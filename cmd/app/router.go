package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/alexedwards/scs/goredisstore"
	"github.com/alexedwards/scs/v2"
	"github.com/dmitrymomot/clientip"
	"github.com/dmitrymomot/go-app-template/pkg/logger"
	"github.com/dmitrymomot/go-app-template/web/templates/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	httprateredis "github.com/go-chi/httprate-redis"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

// initRouter initializes and configures the router for the application.
// It sets up the middleware stack, handles CORS, disables caching in debug mode,
// and registers default error handlers. It also handles serving static files
// from the './web/static' subdirectory.
func initRouter(log *zap.SugaredLogger, redisClient *redis.Client) *chi.Mux {
	r := chi.NewRouter()

	// Middleware stack
	r.Use(
		middleware.Heartbeat("/health"),
		middleware.ThrottleBacklog(httpTrottleLimit, httpTrottleBacklog, httpTrottleTimeout),
		clientip.Middleware(),
		httprate.LimitByRealIP(httpRequestLimit, httpRateLimitWindow), // Limit requests per IP
		httprate.Limit(
			httpRequestLimit,
			httpRateLimitWindow,
			httprate.WithKeyByIP(),
			httprateredis.WithRedisLimitCounter(&httprateredis.Config{
				Client: redisClient,
			}),
		),
		logger.LogRequest(log),
		middleware.Recoverer,
		middleware.CleanPath,
		middleware.StripSlashes,
		middleware.GetHead,
		middleware.Timeout(httpRequestTimeout),
		middleware.SetHeader("X-Content-Type-Options", "nosniff"), // Protection against MIME-sniffing
		middleware.SetHeader("X-Frame-Options", "deny"),           // Protection against clickjacking
		middleware.SetHeader("Server", serverHeader),

		// Basic CORS
		// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
		cors.Handler(cors.Options{
			AllowedOrigins:   corsAllowedOrigins,
			AllowedMethods:   corsAllowedMethods,
			AllowedHeaders:   corsAllowedHeaders,
			AllowCredentials: corsAllowedCredentials,
			MaxAge:           corsMaxAge, // Maximum value not ignored by any of major browsers
		}),

		// TODO: route headers, useful for setting different routers for subdomains
		// For more details, see https://go-chi.io/#/pages/middleware?id=routeheaders
		// middleware.RouteHeaders(),

		// CSRF protection
		// For more details, see https://github.com/gorilla/csrf?tab=readme-ov-file#html-forms
		// csrf.Protect(csrfSecret,
		// 	csrf.RequestHeader("X-CSRF-Token"),
		// 	csrf.CookieName("X-CSRF-Token"),
		// 	csrf.FieldName("_csrf"),
		// 	csrf.SameSite(csrf.SameSiteLaxMode),
		// 	csrf.Secure(appEnv == "production"),
		// 	csrf.TrustedOrigins(corsAllowedOrigins), // Allow cross-domain CSRF use-cases
		// 	csrf.ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 		sendErrorResponse(w, r, http.StatusForbidden, errors.New("CSRF token invalid"))
		// 	})),
		// ),
	)

	// Disable caching
	if disableHTTPCache {
		r.Use(middleware.NoCache)
	}

	// Initialize a new session manager and configure the session lifetime.
	sessionManager := scs.New()
	sessionManager.Lifetime = sessionTTL
	sessionManager.Cookie.Name = sessionName
	sessionManager.Cookie.Secure = appEnv == EnvProduction
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Store = goredisstore.NewWithPrefix(redisClient, sessionPrefix)
	r.Use(sessionManager.LoadAndSave)

	// Default error handlers
	r.NotFound(notFoundHandler())
	r.MethodNotAllowed(methodNotAllowedHandler())

	if appDebugMode {
		// Profiler endpoints, only for debug mode
		r.Mount("/debug", middleware.Profiler())
	}

	// Static file serving from '/assets' subdirectory without directory listing.
	if _, err := os.Stat(staticDir); !os.IsNotExist(err) {
		if err := fileServer(r, staticURLPrefix, http.Dir(staticDir), staticCacheTTL); err != nil {
			log.Fatal(err)
		}
	}

	return r
}

// notFoundHandler is a handler for 404 Not Found
func notFoundHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := errors.New("Page not found")
		if isJsonRequest(r) {
			err = errors.New("Endpoint not found")
		}
		sendErrorResponse(
			w, r,
			http.StatusNotFound,
			err,
		)
	}
}

// methodNotAllowedHandler is a handler for 405 Method Not Allowed
func methodNotAllowedHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sendErrorResponse(
			w, r,
			http.StatusMethodNotAllowed,
			errors.New(http.StatusText(http.StatusMethodNotAllowed)),
		)
	}
}

// Predefined http encoder content type
const (
	contentTypeHeader  = "Content-Type"
	contextTypeCharset = "charset=utf-8"
	contentTypeJSON    = "application/json"
	contentTypeHTML    = "text/html"
	contentTypeJSONUTF = contentTypeJSON + "; " + contextTypeCharset
	contentTypeHTMLUTF = contentTypeHTML + "; " + contextTypeCharset
)

// Helper function to check if an error code is valid
func isValidErrorCode(errCode int) bool {
	return errCode >= 400 && errCode < 600
}

// Is request a json request?
func isJsonRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get(contentTypeHeader)), contentTypeJSON)
}

// Helper function to send an error response
func sendErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	if !isValidErrorCode(statusCode) {
		statusCode = http.StatusInternalServerError
	}

	if isJsonRequest(r) {
		w.Header().Set(contentTypeHeader, contentTypeJSONUTF)
		w.WriteHeader(statusCode)
		if err := json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set(contentTypeHeader, contentTypeHTMLUTF)
	w.WriteHeader(statusCode)
	if err := views.ErrorPage(statusCode, err.Error()).Render(r.Context(), w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
