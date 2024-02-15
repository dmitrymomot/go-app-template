package main

import (
	"time"

	"github.com/dmitrymomot/go-env"
	_ "github.com/joho/godotenv/autoload" // Load .env file automatically
)

// Enviroments
const (
	EnvLocal       = "local"
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
	EnvTesting     = "testing"
)

// Predefined environment variables
var (
	// App
	appName      = env.GetString("APP_NAME", "go-app-template")
	appEnv       = env.GetString("APP_ENV", EnvProduction) // local, development, production, testing
	appDebugMode = env.GetBool("APP_DEBUG_MODE", false)
	appLogLevel  = env.GetString("APP_LOG_LEVEL", "info") // debug, info, warn, error

	// Build
	buildTag = env.GetString("COMMIT_HASH", "undefined")

	// DB
	dbConnString   = env.MustString("DATABASE_URL")
	dbMaxOpenConns = env.GetInt("DATABASE_MAX_OPEN_CONNS", 20)
	dbMaxIdleConns = env.GetInt("DATABASE_IDLE_CONNS", 2)

	// HTTP
	httpPort            = env.GetInt("HTTP_PORT", 8080)
	serverHeader        = env.GetString("SERVER_HEADER", appName+"/"+buildTag)
	httpRequestTimeout  = env.GetDuration("HTTP_REQUEST_TIMEOUT", 5*time.Second)
	httpTrottleLimit    = env.GetInt("HTTP_TROTTLE_LIMIT", 1000)
	httpTrottleBacklog  = env.GetInt("HTTP_TROTTLE_BACKLOG", 1000)
	httpTrottleTimeout  = env.GetDuration("HTTP_TROTTLE_TIMEOUT", time.Second)
	httpRequestLimit    = env.GetInt("HTTP_REQUEST_LIMIT", 100)
	httpRateLimitWindow = env.GetDuration("HTTP_RATE_LIMIT_WINDOW", 1*time.Minute)
	// httpBodyLimit = env.GetInt("HTTP_BODY_LIMIT", 4*1024*1024) // 4MB, 4194304 bytes
	httpReadTimeout  = env.GetDuration("HTTP_READ_TIMEOUT", 5*time.Second)
	httpWriteTimeout = env.GetDuration("HTTP_WRITE_TIMEOUT", 10*time.Second)

	// CORS
	corsAllowedOrigins     = env.GetStrings("CORS_ALLOWED_ORIGINS", ",", []string{"*"})
	corsAllowedMethods     = env.GetStrings("CORS_ALLOWED_METHODS", ",", []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"})
	corsAllowedHeaders     = env.GetStrings("CORS_ALLOWED_HEADERS", ",", []string{"*"})
	corsAllowedCredentials = env.GetBool("CORS_ALLOWED_CREDENTIALS", true)
	corsMaxAge             = env.GetInt("CORS_MAX_AGE", 300)

	// CSRF
	csrfSecret = env.GetBytes("CSRF_SECRET", []byte("32-byte-long-auth-key"))

	// Static
	staticDir       = env.GetString("STATIC_DIR", "./web/static")   // Must be a relative path
	staticURLPrefix = env.GetString("STATIC_URL_PREFIX", "/static") // Must start with a slash
	staticCacheTTL  = env.GetDuration("STATIC_CACHE_TTL", time.Hour)

	// Cache
	disableHTTPCache = env.GetBool("DISABLE_HTTP_CACHE", true)

	// Redis
	redisConnString = env.GetString("REDIS_URL", "redis://localhost:6379/0")

	// Session
	sessionName   = env.GetString("SESSION_COOKIE_NAME", "session")
	sessionPrefix = env.GetString("SESSION_PREFIX", "session:")
	sessionTTL    = env.GetDuration("SESSION_TTL", 24*time.Hour)
)
