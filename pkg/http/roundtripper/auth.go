package roundtripper

import (
	"net/http"

	httpctx "github.com/akuity/grpc-gateway-client/pkg/http/context"
)

var (
	_ WrappedRoundTripper = &authorizationHeaderInjector{}
)

type authorizationHeaderInjector struct {
	rt http.RoundTripper
}

func (i *authorizationHeaderInjector) Unwrap() http.RoundTripper {
	return i.rt
}

func (i *authorizationHeaderInjector) RoundTrip(req *http.Request) (*http.Response, error) {
	ctx := req.Context()
	auth, ok := httpctx.GetAuthorizationHeader(ctx)
	if !ok || len(auth) == 0 {
		return i.rt.RoundTrip(req)
	}

	newReq := req.Clone(ctx)
	newReq.Header.Del("Authorization")
	newReq.Header.Set("Authorization", auth)
	return i.rt.RoundTrip(newReq)
}

func ApplyAuthorizationHeaderInjector(client *http.Client) {
	client.Transport = &authorizationHeaderInjector{
		rt: GetRoundTripper(client),
	}
}
