package gateway

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

	"github.com/akuity/grpc-gateway-client/test/gen/healthv1"
)

type healthServer struct {
	healthv1.UnimplementedHealthServer
}

func (s *healthServer) Check(_ context.Context, _ *healthv1.HealthCheckRequest) (*healthv1.HealthCheckResponse, error) {
	return &healthv1.HealthCheckResponse{Status: healthv1.HealthCheckResponse_SERVING}, nil
}

func (s *healthServer) Watch(_ *healthv1.HealthCheckRequest, srv healthv1.Health_WatchServer) error {
	t := time.NewTicker(10 * time.Millisecond)
	for {
		select {
		case <-t.C:
			_ = srv.Send(&healthv1.HealthCheckResponse{Status: healthv1.HealthCheckResponse_SERVING})
		case <-srv.Context().Done():
			return nil
		}
	}
}

type RequestTestSuite struct {
	suite.Suite

	l       *bufconn.Listener
	grpcSrv *grpc.Server
	gwSrv   *httptest.Server
	client  Client
}

func (s *RequestTestSuite) SetupTest() {
	s.l = bufconn.Listen(256 * 1024)
	s.grpcSrv = grpc.NewServer()
	healthv1.RegisterHealthServer(s.grpcSrv, &healthServer{})
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
	sseMarshaller := NewEventStreamMarshaller(marshaller)
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption("application/json", marshaller),
		runtime.WithMarshalerOption("text/event-stream", sseMarshaller),
	)
	s.Require().NoError(healthv1.RegisterHealthHandler(context.TODO(), mux, cc))
	s.gwSrv = httptest.NewServer(mux)
	s.client = NewClient(s.gwSrv.URL)
}

func (s *RequestTestSuite) TestDoRequest() {
	req := s.client.NewRequest(http.MethodGet, "/healthz")
	res, err := DoRequest[healthv1.HealthCheckResponse](context.TODO(), req)
	s.Require().NoError(err)
	s.Require().Equal(healthv1.HealthCheckResponse_SERVING, res.Status)
}

func (s *RequestTestSuite) TestDoStreamingRequest() {
	ctx, cancel := context.WithTimeout(context.TODO(), 300*time.Millisecond)
	defer cancel()

	req := s.client.NewRequest(http.MethodGet, "/healthz/watch")
	resCh, _, err := DoStreamingRequest[healthv1.HealthCheckResponse](ctx, s.client, req)
	s.Require().NoError(err)
	for {
		select {
		case <-ctx.Done():
			return
		case data := <-resCh:
			s.Require().Equal(healthv1.HealthCheckResponse_SERVING, data.Status)
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
