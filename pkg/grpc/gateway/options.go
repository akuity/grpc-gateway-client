package gateway

import "net/http"

type ClientOption func(*client)

func WithHTTPClient(hc *http.Client) ClientOption {
	return func(c *client) {
		c.httpClient = hc
	}
}

func SkipTLSVerify(skip bool) ClientOption {
	return func(c *client) {
		c.skipTLSVerify = skip
	}
}
