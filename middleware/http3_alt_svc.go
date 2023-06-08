package middleware

import (
	"net/http"

	"github.com/skerkour/golibs/httputils"
)

func Http3AltSvc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(httputils.HeaderAltSvc, `h3=":443"; ma=86400, h3-29=":443"; ma=86400`)
		next.ServeHTTP(w, r)
	})
}
