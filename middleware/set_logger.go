package middleware

import (
	"net/http"

	"github.com/skerkour/golibs/slogutil"
	"github.com/skerkour/golibs/ulid"
	"golang.org/x/exp/slog"
)

// SetLogger injects `logger` in the context of requests
func SetLogger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, req *http.Request) {
			ctx := req.Context()
			routeLogger := logger.With(
				slog.Group("http", slog.String("method", req.Method), slog.String("path", req.URL.Path)),
			)

			reqIDContextValue := ctx.Value(RequestIDCtxKey)
			if requestID, ok := reqIDContextValue.(ulid.ULID); ok {
				routeLogger = routeLogger.With(slog.String("request_id", requestID.String()))
			}

			ctx = slogutil.ToCtx(ctx, routeLogger)
			req = req.WithContext(ctx)

			next.ServeHTTP(w, req)
		}
		return http.HandlerFunc(fn)

	}
}
