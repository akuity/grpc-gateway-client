package gateway

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/alevinval/sse/pkg/decoder"
	"github.com/go-resty/resty/v2"
	"google.golang.org/genproto/googleapis/api/httpbody"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/status"
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
		return nil, status.Error(HTTPStatusToCode(res.StatusCode()), errRes.String())
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
		eventDecoder := decoder.New(body)
		for {
			event, err := eventDecoder.Decode()
			if err != nil {
				if errors.Is(err, io.EOF) {
					close(resCh)
					return
				}
				errCh <- err
				return
			}

			var res streamingResponse
			if err := json.Unmarshal([]byte(event.GetData()), &res); err != nil {
				errCh <- fmt.Errorf("unmarshal streaming response: %w", err)
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
		return nil, status.Error(HTTPStatusToCode(res.StatusCode()), errRes.String())
	}
	return &httpbody.HttpBody{
		ContentType: res.Header().Get("Content-Type"),
		Data:        res.Body(),
	}, nil
}

func doHTTPStreamingRequest(ctx context.Context, c Client, req *resty.Request) (any, <-chan error, error) {
	res, err := req.SetContext(ctx).
		SetHeader("Accept", "text/event-stream").
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
		r := bufio.NewReader(body)
		for {
			var data bytes.Buffer
			for {
				line, isPrefix, err := r.ReadLine()
				if err != nil {
					if errors.Is(err, io.EOF) {
						close(resCh)
						return
					}
					errCh <- fmt.Errorf("read line: %w", err)
					return
				}
				if _, err := data.Write(append(line, '\n')); err != nil {
					errCh <- fmt.Errorf("write line: %w", err)
					return
				}
				if !isPrefix {
					break
				}
			}
			resCh <- &httpbody.HttpBody{
				ContentType: contentType,
				Data:        data.Bytes(),
			}
		}
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

	var streamingResp streamingResponse
	if err := json.Unmarshal(data, &streamingResp); err != nil {
		return fmt.Errorf("unmarshal raw response: %w", err)
	}
	rawErrRes, ok := streamingResp[streamingResponseErrorKey]
	if !ok {
		return errors.New(string(data))
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
