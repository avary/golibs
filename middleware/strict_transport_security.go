package middleware

import (
	"net/http"

	"github.com/skerkour/golibs/httputils"
)

// StrictTransportSecurity sets the Strict-Transport-Security header to maxAge
// if maxAge is empty, it's set to 63072000
func StrictTransportSecurity(maxAge *string, includeSubDomains bool) func(next http.Handler) http.Handler {
	maxAgeKey := "max-age="
	headerValue := maxAgeKey + "63072000"

	if maxAge != nil {
		headerValue = maxAgeKey + *maxAge
	}

	if includeSubDomains {
		headerValue += "; includeSubDomains"
	}

	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set(httputils.HeaderStrictTransportSecurity, headerValue)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
