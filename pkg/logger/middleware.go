package logger

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

// LogRequest is a middleware function that logs incoming HTTP requests.
// It takes a *zap.SugaredLogger as input and returns a function that can be used as middleware.
// The returned function wraps the provided http.Handler and logs information about the request.
// It logs the HTTP method, remote address, request path, and duration of the request.
func LogRequest(log *zap.SugaredLogger) func(http.Handler) http.Handler {
	return LogRequestWithSkipper(log, DefaultSkipper)
}

// LogRequestWithSkipper is a middleware function that logs incoming HTTP requests.
// It takes a *zap.SugaredLogger and a Skipper function as input and returns a function that can be used as middleware.
// The returned function wraps the provided http.Handler and logs information about the request.
// It logs the HTTP method, remote address, request path, and duration of the request.
// The Skipper function is used to determine if the request should be logged.
func LogRequestWithSkipper(log *zap.SugaredLogger, skipper func(r *http.Request) bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if skipper(r) {
				next.ServeHTTP(w, r)
				return
			}

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			start := time.Now()
			defer func() {
				finish := time.Since(start)
				log.Debugw("Request",
					"method", r.Method,
					"remote", r.RemoteAddr,
					"path", r.URL.Path,
					"duration", finish,
					"status", ww.Status(),
					"size", fmt.Sprintf("%dB", ww.BytesWritten()),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

// DefaultSkipper is a function that always returns false.
// It can be used as a default value for the skipper argument of LogRequestWithSkipper.
func DefaultSkipper(r *http.Request) bool {
	if r.Method == http.MethodHead ||
		r.Method == http.MethodOptions ||
		r.URL.Path == "/health" ||
		r.URL.Path == "/favicon.ico" ||
		r.URL.Path == "/robots.txt" ||
		strings.HasPrefix(r.URL.Path, "/static") {
		return true
	}
	return false
}
