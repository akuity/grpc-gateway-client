package gateway

import (
	"bytes"
	"testing"

	"github.com/alevinval/sse/pkg/decoder"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/akuity/grpc-gateway-client/test/gen/healthv1"
)

func TestEventStreamMarshaller(t *testing.T) {
	testSets := map[string]struct {
		input    proto.Message
		expected proto.Message
	}{
		"message": {
			input: &healthv1.HealthCheckResponse{
				Status: healthv1.HealthCheckResponse_SERVING,
			},
			expected: &healthv1.HealthCheckResponse{},
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			jm := &runtime.JSONPb{}
			m := NewEventStreamMarshaller(jm)
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
