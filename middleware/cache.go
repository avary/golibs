package middleware

import (
	"net/http"

	"github.com/skerkour/golibs/httputils"
)

func NoCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set(httputils.HeaderCacheControl, httputils.CacheControlNoCache)
		w.Header().Set(httputils.HeaderExpires, "0") // Proxies

		next.ServeHTTP(w, r)
	})
}
