package middleware

import (
	"fmt"
	"net/http"

	"github.com/skerkour/golibs/httputils"
)

func Http3AltSvc(next http.Handler, port *string) http.Handler {
	portStr := "443"
	if port != nil {
		portStr = *port
	}
	headerValue := fmt.Sprintf(`h3=":%s"; ma=86400, h3-29=":%s"; ma=86400`, portStr, portStr)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(httputils.HeaderAltSvc, headerValue)
		next.ServeHTTP(w, r)
	})
}
