package gateway

import (
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type ClientOption func(*client)

func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *client) {
		c.httpClient = hc
	}
}

func WithMarshaller(m *runtime.JSONPb) ClientOption {
	return func(c *client) {
		c.marshaller = m
	}
}

func SkipTLSVerify(skip bool) ClientOption {
	return func(c *client) {
		c.skipTLSVerify = skip
	}
}
