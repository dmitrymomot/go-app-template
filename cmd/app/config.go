package main

import (
	"github.com/dmitrymomot/go-env"
)

var (
	// App
	appName      = env.GetString("APP_NAME", "go-app-template")
	appEnv       = env.GetString("APP_ENV", "local") // local, development, production
	appDebugMode = env.GetBool("APP_DEBUG_MODE", false)
	appLogLevel  = env.GetString("APP_LOG_LEVEL", "info") // debug, info, warn, error

	// Build
	buildTag = env.GetString("COMMIT_HASH", "undefined")

	// DB
	dbConnString = env.MustString("DATABASE_URL")

	// HTTP
	httpPort = env.GetInt("HTTP_PORT", 8080)
	// serverHeader  = env.GetString("SERVER_HEADER", appName)
	// httpBodyLimit = env.GetInt("HTTP_BODY_LIMIT", 4*1024*1024) // 4MB, 4194304 bytes
	// readTimeout   = env.GetDuration("READ_TIMEOUT", 5*time.Second)
	// writeTimeout  = env.GetDuration("WRITE_TIMEOUT", 10*time.Second)

	// Static
	// staticDir       = env.GetString("STATIC_DIR", "./web/static")
	// staticURLPrefix = env.GetString("STATIC_URL_PREFIX", "static")
	// staticCacheTTL  = env.GetInt("STATIC_CACHE_TTL", 3600)
)
