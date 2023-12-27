package server

import (
	"bytes"
	"context"
	"encoding/base64"
	"time"

	"github.com/bufbuild/protoyaml-go"
	"google.golang.org/genproto/googleapis/api/httpbody"

	"github.com/omrikiei/grpc-gateway-client/internal/test/gen/testv1"
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

func (s *testServiceServer) TrackInvitation(_ *testv1.TrackInvitationRequest, srv testv1.TestService_TrackInvitationServer) error {
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
