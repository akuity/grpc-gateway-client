package context

import (
	"context"
	"strings"

	ctxutil "github.com/omrikiei/grpc-gateway-client/pkg/context"
)

type authorizationHeaderKey struct {
	/* explicitly empty */
}

func SetAuthorizationHeader(ctx context.Context, scheme, params string) context.Context {
	v := strings.TrimSpace(strings.Join([]string{scheme, params}, " "))
	return ctxutil.Set(ctx, authorizationHeaderKey{}, v)
}

func GetAuthorizationHeader(ctx context.Context) (string, bool) {
	return ctxutil.Get[authorizationHeaderKey, string](ctx, authorizationHeaderKey{})
}
