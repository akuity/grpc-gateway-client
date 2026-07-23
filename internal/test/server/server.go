package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"strings"
	"time"

	"github.com/bufbuild/protoyaml-go"
	"google.golang.org/genproto/googleapis/api/httpbody"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/akuity/grpc-gateway-client/internal/assets"
	"github.com/akuity/grpc-gateway-client/internal/test/gen/testv1"

	_ "embed"
)

type testServiceServer struct {
	testv1.UnimplementedTestServiceServer
}

func NewTestServer() testv1.TestServiceServer {
	return &testServiceServer{}
}

func (s *testServiceServer) ListInvitations(_ context.Context, req *testv1.ListInvitationsRequest) (*testv1.ListInvitationsResponse, error) {
	return &testv1.ListInvitationsResponse{
		Invitations: []*testv1.Invitation{
			{
				Id:     "test-invitation",
				Labels: req.GetQuery().GetLabels(),
			},
		},
	}, nil
}

func (s *testServiceServer) SendInvitation(_ context.Context, req *testv1.SendInvitationRequest) (*testv1.SendInvitationResponse, error) {
	return &testv1.SendInvitationResponse{
		Id: base64.StdEncoding.EncodeToString([]byte(req.Email)),
	}, nil
}

func (s *testServiceServer) TrackInvitation(req *testv1.TrackInvitationRequest, srv testv1.TestService_TrackInvitationServer) error {
	// Used by tests to exercise a server stream that fails *before* its first
	// message. grpc-gateway writes this through the SSE marshaller too, so the
	// HTTP error body is framed as `data: {"error": ...}`.
	if req.GetId() == "fail-before-events" {
		return status.Error(codes.NotFound, "nope")
	}
	// Used by tests to exercise events whose SSE data line far exceeds
	// bufio.MaxScanTokenSize (64KiB): a line-limited decoder either errors or,
	// worse, silently truncates the stream.
	if req.GetId() == "large-events" {
		large := strings.Repeat("x", 300*1024)
		for i := 0; i < 3; i++ {
			if err := srv.Send(&testv1.TrackInvitationResponse{
				Type:    testv1.EventType_EVENT_TYPE_SEEN,
				Message: large,
			}); err != nil {
				return err
			}
		}
		return nil
	}
	eventTypes := []testv1.EventType{
		testv1.EventType_EVENT_TYPE_SEEN,
		testv1.EventType_EVENT_TYPE_ACCEPTED,
	}
	for _, et := range eventTypes {
		_ = srv.Send(&testv1.TrackInvitationResponse{
			Type: et,
		})
		time.Sleep(100 * time.Millisecond)
	}
	// Used by tests to exercise a server stream that fails *after* emitting one
	// or more results, which grpc-gateway frames as a terminal {"error": ...}
	// SSE event.
	if req.GetId() == "fail-after-events" {
		return status.Error(codes.Internal, "boom")
	}
	return nil
}

func (s *testServiceServer) DownloadInvitations(req *testv1.DownloadInvitationsRequest, srv testv1.TestService_DownloadInvitationsServer) error {
	invitations := []*testv1.Invitation{
		{
			Id: "test-1",
		},
		{
			Id: "test-2",
		},
	}
	for _, invitation := range invitations {
		var buf bytes.Buffer
		data, err := protoyaml.Marshal(invitation)
		if err != nil {
			return err
		}
		_, _ = buf.WriteString("---\n")
		_, _ = buf.Write(data)
		_ = srv.Send(&httpbody.HttpBody{
			ContentType: "application/yaml",
			Data:        buf.Bytes(),
		})
		time.Sleep(100 * time.Millisecond)
	}
	return nil
}

func (s *testServiceServer) DownloadLargeFile(
	req *testv1.DownloadLargeFileRequest,
	srv testv1.TestService_DownloadLargeFileServer,
) error {
	return srv.Send(&httpbody.HttpBody{
		ContentType: "text/plain",
		Data:        []byte(assets.LargeFile),
	})
}
