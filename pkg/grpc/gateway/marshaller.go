package gateway

import (
	"bytes"
	"errors"
	"io"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func NewEventStreamMarshaller(marshaller *runtime.JSONPb) runtime.Marshaler {
	return &eventStreamMarshaller{
		marshaller: marshaller,
	}
}

type eventStreamMarshaller struct {
	marshaller *runtime.JSONPb
}

func (e *eventStreamMarshaller) encode(v interface{}) ([]byte, error) {
	data, err := e.marshaller.Marshal(v)
	if err != nil {
		return nil, err
	}
	var b bytes.Buffer
	b.WriteString("data: ")
	b.Write(data)
	b.WriteRune('\n')
	return b.Bytes(), nil
}

func (e *eventStreamMarshaller) Marshal(v interface{}) ([]byte, error) {
	return e.encode(v)
}

func (e *eventStreamMarshaller) Unmarshal(_ []byte, _ interface{}) error {
	return errors.New("unsupported method: Unmarshal")
}

func (e *eventStreamMarshaller) Decode(_ interface{}) error {
	return errors.New("unsupported method: Decode")
}

func (e *eventStreamMarshaller) NewDecoder(_ io.Reader) runtime.Decoder {
	return e
}

func (e *eventStreamMarshaller) NewEncoder(w io.Writer) runtime.Encoder {
	return runtime.EncoderFunc(func(v interface{}) error {
		data, err := e.encode(v)
		if err != nil {
			return err
		}
		_, err = w.Write(data)
		return err
	})
}

func (e *eventStreamMarshaller) ContentType(_ interface{}) string {
	return "text/event-stream"
}
