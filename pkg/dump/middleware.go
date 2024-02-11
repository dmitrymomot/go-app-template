package dump

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"go.uber.org/zap"
)

// DumpRequest is a middleware that dumps the request body, headers and query params
// to the std log. It is useful for debugging purposes only.
// !!! Do not use it in any environment other than local.
func DumpRequest(log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Dump the request body
			var body []byte
			if r.Body != nil {
				body, _ = io.ReadAll(r.Body)
			}
			// return the courtesy of reading
			r.Body = io.NopCloser(strings.NewReader(string(body)))

			// Dump the request headers
			headers := make(map[string]interface{})
			for k, v := range r.Header {
				headers[k] = v
			}

			// Dump the request query params
			query := make(map[string]interface{})
			for k, v := range r.URL.Query() {
				query[k] = v
			}

			// Log the request
			if log != nil {
				log.Infow("[REQ_DEBUG] Request dump",
					"method", r.Method,
					"url", r.URL.String(),
					"headers", headers,
					"query", query,
					"body", string(body),
				)
			} else {
				// If the logger is not provided, use the std log instead.
				b, _ := json.MarshalIndent(map[string]interface{}{
					"method":  r.Method,
					"url":     r.URL.String(),
					"headers": headers,
					"query":   query,
					"body":    string(body),
				}, "", "  ")
				fmt.Println("[REQ_DEBUG] Request dump:\n", string(b))
			}

			// Call the next middleware
			next.ServeHTTP(w, r)
		})
	}
}
