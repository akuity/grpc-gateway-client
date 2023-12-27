package gateway

import (
	"crypto/tls"
	"net/http"

	"google.golang.org/grpc/codes"

	"github.com/omrikiei/grpc-gateway-client/pkg/http/roundtripper"
)

func setSkipTLSVerify(hc *http.Client, skip bool) {
	rt := roundtripper.GetRoundTripper(hc)
	for {
		ht, ok := rt.(*http.Transport)
		if ok {
			if ht.TLSClientConfig == nil {
				ht.TLSClientConfig = &tls.Config{}
			}
			ht.TLSClientConfig.InsecureSkipVerify = skip
			break
		}

		wrapped, ok := rt.(roundtripper.WrappedRoundTripper)
		if !ok {
			break
		}
		rt = wrapped.Unwrap()
	}
}

// HTTPStatusToCode converts HTTP status code to gRPC error code
// https://github.com/grpc/grpc/blob/9d2a1a3d1aba56d94c56d2e01cf2511d1a082445/doc/http-grpc-status-mapping.md
func HTTPStatusToCode(code int) codes.Code {
	switch code {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.Internal
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.Unimplemented
	case http.StatusTooManyRequests:
		return codes.Unavailable
	case http.StatusBadGateway:
		return codes.Unavailable
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	case http.StatusGatewayTimeout:
		return codes.Unavailable
	default:
		return codes.Unknown
	}
}
