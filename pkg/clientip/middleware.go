package clientip

import (
	"net/http"
)

// Middleware returns a middleware function that sets the remote address of the incoming request
// to the value obtained from the LookupFromRequest function. It then calls the next handler in the chain.
func Middleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			if rip := LookupFromRequest(r); rip != "" {
				r.RemoteAddr = rip
			}
			h.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}
