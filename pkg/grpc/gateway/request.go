package gateway

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/go-resty/resty/v2"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
)

type streamingResponse map[string]json.RawMessage

const (
	streamingResponseResultKey = "result"
	streamingResponseErrorKey  = "error"
)

func DoRequest[T any](ctx context.Context, req *resty.Request) (*T, error) {
	var resBody T
	if _, ok := any(&resBody).(*httpbody.HttpBody); ok {
		res, err := doHTTPRequest(ctx, req)
		if err != nil {
			return nil, err
		}
		return res.(*T), nil
	}

	res, err := req.SetContext(ctx).
		SetResult(&resBody).
		SetError(&rpcstatus.Status{}).
		Send()
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	if res.IsError() {
		errRes, ok := res.Error().(*rpcstatus.Status)
		if !ok {
			return nil, fmt.Errorf("cast error response: %s", res.String())
		}
		if err := status.ErrorProto(errRes); err != nil {
			return nil, err
		}
		return nil, status.Error(HTTPStatusToCode(res.StatusCode()), res.String())
	}

	data, ok := res.Result().(*T)
	if !ok {
		return nil, fmt.Errorf("cast response: %s", res.String())
	}
	return data, nil
}

func DoStreamingRequest[T any](ctx context.Context, c Client, req *resty.Request) (<-chan *T, <-chan error, error) {
	var resBody T
	if _, ok := any(&resBody).(*httpbody.HttpBody); ok {
		resCh, errCh, err := doHTTPStreamingRequest(ctx, c, req)
		if err != nil {
			return nil, nil, err
		}
		return resCh.(chan *T), errCh, nil
	}

	rawRes, err := req.SetContext(ctx).
		SetHeader("Accept", "text/event-stream").
		SetHeader("Cache-Control", "no-cache").
		SetHeader("Connection", "keep-alive").
		SetDoNotParseResponse(true).
		Send()
	if err != nil {
		return nil, nil, fmt.Errorf("send request: %w", err)
	}
	if rawRes.IsError() {
		return nil, nil, wrapStreamingResponseError(c, rawRes)
	}

	resCh := make(chan *T)
	errCh := make(chan error)

	go func() {
		body := rawRes.RawBody()
		defer func() { _ = body.Close() }()
		eventDecoder := newSSEEventDecoder(body)
		for {
			payload, err := eventDecoder.next()
			if err != nil {
				if errors.Is(err, io.EOF) {
					close(resCh)
					return
				}
				errCh <- err
				return
			}

			var res streamingResponse
			if err := json.Unmarshal([]byte(payload), &res); err != nil {
				errCh <- fmt.Errorf("unmarshal streaming response: %w", err)
				return
			}
			// A server-streaming RPC that fails after emitting one or more
			// results terminates the stream with a final {"error": <google.rpc.Status>}
			// event. Surface it instead of skipping it; otherwise a truncated
			// stream is indistinguishable from a successful one and the caller
			// silently receives partial data.
			if rawErr, ok := res[streamingResponseErrorKey]; ok {
				var errResp rpcstatus.Status
				if err := c.Unmarshal(rawErr, &errResp); err != nil {
					errCh <- fmt.Errorf("unmarshal streaming error: %w", err)
					return
				}
				if err := status.ErrorProto(&errResp); err != nil {
					errCh <- err
				} else {
					// Defensive: an error event carrying an OK/empty status must
					// still not read as a successful completion.
					errCh <- errors.New("streaming response terminated with an error")
				}
				return
			}
			rawResult, ok := res[streamingResponseResultKey]
			if !ok {
				continue
			}

			var data T
			if err := c.Unmarshal(rawResult, &data); err != nil {
				errCh <- err
				return
			}
			resCh <- &data
		}
	}()
	return resCh, errCh, nil
}

// sseEventDecoder reads Server-Sent Events without any line-length limit and
// propagates transport errors instead of flattening them into EOF. The
// alevinval/sse decoder it replaces capped lines at bufio.MaxScanTokenSize
// (64KiB) and returned io.EOF for every scanner failure, so an oversized or
// mid-line-interrupted event silently ended the stream as if it had
// completed — the caller received partial data with no error.
type sseEventDecoder struct {
	r *bufio.Reader
	// skipLF records that the previous line was terminated by a bare CR whose
	// following byte has not been read yet. If that byte turns out to be LF it
	// is the second half of a CRLF pair and must be swallowed, per the SSE
	// parsing algorithm. Tracking this across reads (instead of peeking ahead
	// synchronously) lets a CR-terminated line be delivered immediately even
	// when the CR is the last byte the server has flushed so far.
	skipLF bool
}

func newSSEEventDecoder(r io.Reader) *sseEventDecoder {
	return &sseEventDecoder{r: bufio.NewReader(r)}
}

// next returns the data payload of the next event. io.EOF marks a clean
// end-of-stream at an event boundary; an EOF that interrupts a partially
// read event surfaces as io.ErrUnexpectedEOF so truncation is never mistaken
// for successful completion. Non-data fields (event, id, retry) and comment
// lines are ignored: the grpc-gateway SSE framing carries everything in
// `data:` lines.
func (d *sseEventDecoder) next() (string, error) {
	var data bytes.Buffer
	dataSeen := false
	for {
		line, err := d.readLine()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if line != "" || dataSeen {
					return "", io.ErrUnexpectedEOF
				}
				return "", io.EOF
			}
			return "", err
		}

		if line == "" {
			if dataSeen {
				return data.String(), nil
			}
			continue
		}

		field, value, _ := strings.Cut(line, ":")
		if field == "" {
			// Comment line (":heartbeat" and friends).
			continue
		}
		if field == "data" {
			// Per the SSE spec, multiple data lines of one event are joined
			// with a newline, and a single leading space is trimmed.
			if dataSeen {
				data.WriteByte('\n')
			}
			data.WriteString(strings.TrimPrefix(value, " "))
			dataSeen = true
		}
	}
}

// readLine reads one line terminated by LF, CRLF, or bare CR — the three
// end-of-line sequences the SSE grammar permits. The terminator is consumed
// and excluded from the returned line. A non-nil error means the line was
// never terminated; for io.EOF the bytes read so far are returned alongside
// it so the caller can tell a clean boundary from a truncated line.
func (d *sseEventDecoder) readLine() (string, error) {
	var line []byte
	for {
		b, err := d.r.ReadByte()
		if err != nil {
			return string(line), err
		}
		if d.skipLF {
			d.skipLF = false
			if b == '\n' {
				continue
			}
		}
		switch b {
		case '\n':
			return string(line), nil
		case '\r':
			d.skipLF = true
			return string(line), nil
		default:
			line = append(line, b)
		}
	}
}

func doHTTPRequest(ctx context.Context, req *resty.Request) (any, error) {
	res, err := req.SetContext(ctx).
		SetError(&rpcstatus.Status{}).
		Send()
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}
	if res.IsError() {
		errRes, ok := res.Error().(*rpcstatus.Status)
		if !ok {
			return nil, fmt.Errorf("cast error response: %s", res.String())
		}
		if err := status.ErrorProto(errRes); err != nil {
			return nil, err
		}
		return nil, status.Error(HTTPStatusToCode(res.StatusCode()), res.String())
	}
	return &httpbody.HttpBody{
		ContentType: res.Header().Get("Content-Type"),
		Data:        res.Body(),
	}, nil
}

func doHTTPStreamingRequest(ctx context.Context, c Client, req *resty.Request) (any, <-chan error, error) {
	res, err := req.SetContext(ctx).
		SetHeader("Cache-Control", "no-cache").
		SetHeader("Connection", "keep-alive").
		SetDoNotParseResponse(true).
		Send()
	if err != nil {
		return nil, nil, fmt.Errorf("send request: %w", err)
	}
	if res.IsError() {
		return nil, nil, wrapStreamingResponseError(c, res)
	}

	resCh := make(chan *httpbody.HttpBody)
	errCh := make(chan error)
	go func() {
		contentType := res.Header().Get("Content-Type")
		body := res.RawBody()
		defer func() { _ = body.Close() }()

		var data bytes.Buffer
		if _, err := io.Copy(&data, body); err != nil {
			errCh <- fmt.Errorf("copy body: %w", err)
			return
		}
		resCh <- &httpbody.HttpBody{
			ContentType: contentType,
			Data:        data.Bytes(),
		}
		close(resCh)
	}()
	return resCh, errCh, nil
}

func wrapStreamingResponseError(c Client, resp *resty.Response) error {
	body := resp.RawBody()
	defer func() { _ = body.Close() }()
	data, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("read error response body: %w", err)
	}

	// A streaming handler that fails before its first message emits the error
	// through the SSE marshaller, so the body is framed as `data: {"error":{...}}`.
	// Strip that framing so the JSON underneath parses. This is a no-op for
	// non-streaming (plain JSON) error bodies such as routing 404s, which start
	// with '{' rather than the "data:" prefix.
	data = bytes.TrimSpace(data)
	data = bytes.TrimSpace(bytes.TrimPrefix(data, []byte("data:")))

	var streamingResp streamingResponse
	if err := json.Unmarshal(data, &streamingResp); err != nil {
		return fmt.Errorf("unmarshal raw response: %w", err)
	}
	rawErrRes, ok := streamingResp[streamingResponseErrorKey]
	if !ok {
		var statusResp rpcstatus.Status
		if err := protojson.Unmarshal(data, &statusResp); err != nil {
			return errors.New(string(data))
		}
		return status.FromProto(&statusResp).Err()
	}
	var errResp rpcstatus.Status
	if err := c.Unmarshal(rawErrRes, &errResp); err != nil {
		return fmt.Errorf("unmarshal error response: %w", err)
	}
	if err := status.ErrorProto(&errResp); err != nil {
		return err
	}
	return status.Error(HTTPStatusToCode(resp.StatusCode()), resp.String())
}
