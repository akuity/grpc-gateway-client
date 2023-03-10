package gateway_test

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"

	"github.com/akuity/grpc-gateway-client/internal/test/gen/testv1"
	"github.com/akuity/grpc-gateway-client/internal/test/server"
	"github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
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

func (s *RequestTestSuite) TearDownTest() {
	s.gwSrv.Close()
	s.grpcSrv.Stop()
	s.l.Close()
}

func TestRequestTestSuite(t *testing.T) {
	suite.Run(t, &RequestTestSuite{})
}
