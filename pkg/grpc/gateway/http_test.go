package gateway

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/omrikiei/grpc-gateway-client/pkg/http/roundtripper"
)

var (
	_ roundtripper.WrappedRoundTripper = &wrappedTransport{}
	_ http.RoundTripper                = &unwrappableTransport{}
)

type wrappedTransport struct {
	rt http.RoundTripper
}

func (t *wrappedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.rt.RoundTrip(req)
}

func (t *wrappedTransport) Unwrap() http.RoundTripper {
	return t.rt
}

type unwrappableTransport struct {
	rt http.RoundTripper
}

func (t unwrappableTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.rt.RoundTrip(req)
}

func Test_setSkipTLSVerify(t *testing.T) {
	testSets := map[string]struct {
		newClientFunc func() *http.Client
		errExpected   bool
	}{
		"client with default transport": {
			newClientFunc: func() *http.Client {
				return &http.Client{
					Transport: http.DefaultTransport.(*http.Transport).Clone(),
				}
			},
			errExpected: false,
		},
		"client with wrapped transport": {
			newClientFunc: func() *http.Client {
				return &http.Client{
					Transport: &wrappedTransport{
						rt: http.DefaultTransport.(*http.Transport).Clone(),
					},
				}
			},
			errExpected: false,
		},
		"client with unwrappable transport": {
			newClientFunc: func() *http.Client {
				return &http.Client{
					Transport: &unwrappableTransport{
						rt: http.DefaultTransport.(*http.Transport).Clone(),
					},
				}
			},
			errExpected: true,
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusNoContent)
			}))

			hc := ts.newClientFunc()
			setSkipTLSVerify(hc, true)
			res, err := hc.Get(srv.URL)
			if ts.errExpected {
				require.Error(t, err)
			} else {
				require.NotNil(t, res)
				require.Equal(t, http.StatusNoContent, res.StatusCode)
			}
		})
	}
}
