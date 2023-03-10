package gateway

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Client interface {
	NewRequest(method, url string) *resty.Request

	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
}

type client struct {
	httpClient    *http.Client
	skipTLSVerify bool
	rc            *resty.Client

	marshaller *runtime.JSONPb
}

func NewClient(baseURL string, opts ...ClientOption) Client {
	c := &client{}
	for _, opt := range opts {
		opt(c)
	}

	if c.httpClient == nil {
		c.httpClient = &http.Client{}
	}
	if c.httpClient.Transport == nil {
		c.httpClient.Transport = http.DefaultTransport.(*http.Transport).Clone()
	}
	setSkipTLSVerify(c.httpClient, c.skipTLSVerify)

	if c.marshaller == nil {
		c.marshaller = &runtime.JSONPb{}
	}
	c.rc = resty.NewWithClient(c.httpClient).SetBaseURL(baseURL)
	c.rc.JSONMarshal = c.marshaller.Marshal
	c.rc.JSONUnmarshal = c.marshaller.Unmarshal
	return c
}

func (c *client) NewRequest(method string, url string) *resty.Request {
	req := c.rc.NewRequest()
	req.Method = method
	req.URL = url
	return req
}

func (c *client) Marshal(v interface{}) ([]byte, error) {
	return c.marshaller.Marshal(v)
}

func (c *client) Unmarshal(data []byte, v interface{}) error {
	return c.marshaller.Unmarshal(data, v)
}
