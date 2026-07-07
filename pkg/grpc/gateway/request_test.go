package gateway_test

import (
	"bytes"
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/bufbuild/protoyaml-go"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/protobuf/proto"

	"github.com/akuity/grpc-gateway-client/internal/assets"
	"github.com/akuity/grpc-gateway-client/internal/test/gen/testv1"
	"github.com/akuity/grpc-gateway-client/internal/test/server"
	"github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"

	_ "embed"
)

type RequestTestSuite struct {
	suite.Suite

	l       *bufconn.Listener
	grpcSrv *grpc.Server
	gwSrv   *httptest.Server
	client  gateway.Client
}

func (s *RequestTestSuite) SetupTest() {
	s.l = bufconn.Listen(256 * 1024)
	s.grpcSrv = grpc.NewServer()
	testv1.RegisterTestServiceServer(s.grpcSrv, server.NewTestServer())
	go func() {
		_ = s.grpcSrv.Serve(s.l)
	}()

	cc, err := grpc.Dial("",
		grpc.WithContextDialer(func(_ context.Context, _ string) (net.Conn, error) {
			return s.l.Dial()
		}),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	s.Require().NoError(err)

	marshaller := &runtime.JSONPb{}
	sseMarshaller := gateway.NewEventStreamMarshaller(marshaller)
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption("application/json", marshaller),
		runtime.WithMarshalerOption("text/event-stream", sseMarshaller),
	)
	s.Require().NoError(testv1.RegisterTestServiceHandler(context.TODO(), mux, cc))
	s.gwSrv = httptest.NewServer(mux)
	s.client = gateway.NewClient(s.gwSrv.URL)
}

func (s *RequestTestSuite) TestDoRequest() {
	req := s.client.NewRequest(http.MethodPost, "/invitation").
		SetBody(&testv1.SendInvitationRequest{
			Email: "test@test.com",
		})
	res, err := gateway.DoRequest[testv1.SendInvitationResponse](context.TODO(), req)
	s.Require().NoError(err)
	s.Require().NotEmpty(res.GetId())
}

func (s *RequestTestSuite) TestDoStreamingRequest() {
	ctx, cancel := context.WithTimeout(context.TODO(), 300*time.Millisecond)
	defer cancel()

	req := s.client.NewRequest(http.MethodGet, "/invitation/some-id")
	resCh, _, err := gateway.DoStreamingRequest[testv1.TrackInvitationResponse](ctx, s.client, req)
	s.Require().NoError(err)
	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-resCh:
			if !ok {
				return
			}
			s.Require().True(testv1.EventType_EVENT_TYPE_UNDEFINED != data.GetType())
		}
	}
}

func (s *RequestTestSuite) TestDoStreamingRequest_ErrorAfterResults() {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
	defer cancel()

	// TrackInvitation emits two results and then fails; grpc-gateway terminates
	// the stream with a final {"error": ...} SSE event. That error must surface
	// on errCh rather than being silently dropped, otherwise a truncated stream
	// is indistinguishable from a successful one.
	req := s.client.NewRequest(http.MethodGet, "/invitation/fail-after-events")
	resCh, errCh, err := gateway.DoStreamingRequest[testv1.TrackInvitationResponse](ctx, s.client, req)
	s.Require().NoError(err) // the stream starts 200 OK; the failure arrives mid-stream

	var results int
	var streamErr error
read:
	for {
		select {
		case <-ctx.Done():
			s.FailNow("timed out; the terminal error was never delivered on errCh")
		case e := <-errCh:
			streamErr = e
			break read
		case _, ok := <-resCh:
			if !ok {
				break read // clean EOF; before the fix this is the (incorrect) path
			}
			results++
		}
	}

	s.Require().Equal(2, results)
	s.Require().Error(streamErr)
	stat, ok := status.FromError(streamErr)
	s.Require().True(ok)
	s.Require().Equal(codes.Internal, stat.Code())
}

func (s *RequestTestSuite) TestDownloadRequest() {
	ctx, cancel := context.WithTimeout(context.TODO(), 500*time.Millisecond)
	defer cancel()
	req := s.client.NewRequest(http.MethodGet, "/download-invitations")
	resCh, errCh, err := gateway.DoStreamingRequest[httpbody.HttpBody](ctx, s.client, req)
	s.Require().NoError(err)
	var buf bytes.Buffer

read:
	for {
		select {
		case <-ctx.Done():
			break read
		case err := <-errCh:
			s.Require().NoError(err)
		case data, ok := <-resCh:
			if !ok {
				break read
			}
			buf.Write(data.GetData())
		}
	}

	var actual []*testv1.Invitation
	docs := strings.Split(buf.String(), "---\n")
	for _, doc := range docs {
		data := strings.TrimSpace(doc)
		if data == "" {
			continue
		}

		invitation := &testv1.Invitation{}
		s.Require().NoError(protoyaml.Unmarshal([]byte(data), invitation))
		actual = append(actual, invitation)
	}
	expected := []*testv1.Invitation{
		{
			Id: "test-1",
		},
		{
			Id: "test-2",
		},
	}
	s.Require().Len(actual, len(expected))
	for idx := range actual {
		s.Require().True(proto.Equal(expected[idx], actual[idx]))
	}
}

func (s *RequestTestSuite) TestDownloadLargeFileRequest() {
	ctx, cancel := context.WithTimeout(context.TODO(), 100*time.Millisecond)
	defer cancel()
	req := s.client.NewRequest(http.MethodGet, "/download-large-file")
	resCh, errCh, err := gateway.DoStreamingRequest[httpbody.HttpBody](ctx, s.client, req)
	s.Require().NoError(err)
	var buf bytes.Buffer

read:
	for {
		select {
		case <-ctx.Done():
			break read
		case err := <-errCh:
			s.Require().NoError(err)
		case data, ok := <-resCh:
			if !ok {
				break read
			}
			buf.Write(data.GetData())
		}
	}

	require.NoError(s.T(), ctx.Err())
	require.Equal(s.T(), strings.TrimSpace(assets.LargeFile), strings.TrimSpace(buf.String()))
}

func (s *RequestTestSuite) TestDownloadRequest_Error() {
	ctx, cancel := context.WithTimeout(context.TODO(), 500*time.Millisecond)
	defer cancel()
	req := s.client.NewRequest(http.MethodGet, "/download-invitations").
		SetQueryParams(map[string]string{
			"type": "EVENT_TYPE_UNKNOWN",
		})
	_, _, err := gateway.DoStreamingRequest[httpbody.HttpBody](ctx, s.client, req)
	s.Require().Error(err)
	stat, ok := status.FromError(err)
	s.Require().True(ok)
	s.Require().Equal(codes.InvalidArgument, stat.Code())
}

func (s *RequestTestSuite) TearDownTest() {
	s.gwSrv.Close()
	s.grpcSrv.Stop()
	s.l.Close()
}

func TestRequestTestSuite(t *testing.T) {
	suite.Run(t, &RequestTestSuite{})
}
