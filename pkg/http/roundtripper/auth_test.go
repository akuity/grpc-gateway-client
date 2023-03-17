package roundtripper

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	httpctx "github.com/akuity/grpc-gateway-client/pkg/http/context"
)

func TestAuthorizationInjector(t *testing.T) {
	testSets := map[string]struct {
		newContextFunc func() context.Context
		expected       string
	}{
		"context with credential": {
			newContextFunc: func() context.Context {
				return httpctx.SetAuthorizationHeader(context.TODO(), "Bearer", "token")
			},
			expected: "Bearer token",
		},
		"context with empty credential": {
			newContextFunc: func() context.Context {
				return httpctx.SetAuthorizationHeader(context.TODO(), "", "")
			},
			expected: "", // empty string
		},
		"context without credential": {
			newContextFunc: context.TODO,
			expected:       "", // empty string
		},
	}
	for name, ts := range testSets {
		t.Run(name, func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				_, _ = w.Write([]byte(req.Header.Get("Authorization")))
			}))
			defer srv.Close()

			c := &http.Client{}
			ApplyAuthorizationHeaderInjector(c)

			req, err := http.NewRequestWithContext(ts.newContextFunc(), http.MethodGet, srv.URL, nil)
			require.NoError(t, err)
			res, err := c.Do(req)
			require.NoError(t, err)
			actual, err := io.ReadAll(res.Body)
			require.NoError(t, err)
			require.Equal(t, ts.expected, string(actual))
		})
	}
}
