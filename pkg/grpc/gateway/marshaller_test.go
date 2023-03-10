package gateway_test

import (
	"bytes"
	"testing"

	"github.com/alevinval/sse/pkg/decoder"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/akuity/grpc-gateway-client/internal/test/gen/testv1"
	"github.com/akuity/grpc-gateway-client/pkg/grpc/gateway"
)

func TestEventStreamMarshaller(t *testing.T) {
	testSets := map[string]struct {
		input    proto.Message
		expected proto.Message
	}{
		"message": {
			input: &testv1.SendInvitationResponse{
				Id: "some-id",
			},
			expected: &testv1.SendInvitationResponse{
				Id: "some-id",
			},
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			jm := &runtime.JSONPb{}
			m := gateway.NewEventStreamMarshaller(jm)
			data, err := m.Marshal(ts.input)
			require.NoError(t, err)

			buf := bytes.NewBuffer(data)
			buf.WriteRune('\n') // Add empty line to the end
			dec := decoder.New(buf)
			event, err := dec.Decode()
			require.NoError(t, err)
			require.NoError(t, jm.Unmarshal([]byte(event.GetData()), ts.expected))
			require.True(t, proto.Equal(ts.input, ts.expected))
		})
	}
}
