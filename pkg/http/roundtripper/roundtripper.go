package roundtripper

import "net/http"

type WrappedRoundTripper interface {
	http.RoundTripper
	Unwrap() http.RoundTripper
}

func GetRoundTripper(c *http.Client) http.RoundTripper {
	rt := c.Transport
	if rt != nil {
		return rt
	}
	return http.DefaultTransport.(*http.Transport).Clone()
}
