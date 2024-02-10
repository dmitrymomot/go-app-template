package main

import "github.com/dmitrymomot/go-env"

var (
	// App
	appName      = env.GetString("APP_NAME", "go-app-template")
	appEnv       = env.GetString("APP_ENV", "development") // local, development, production
	appDebugMode = env.GetBool("APP_DEBUG_MODE", false)
	appLogLevel  = env.GetString("APP_LOG_LEVEL", "info") // debug, info, warn, error

	// Build
	buildTag = env.GetString("COMMIT_HASH", "undefined")

	// DB
	dbConnString    = env.MustString("DATABASE_URL")
	migrationsDir   = env.GetString("DATABASE_MIGRATIONS_DIR", "./db/sql/migrations")
	migrationsTable = env.GetString("DATABASE_MIGRATIONS_TABLE", "migrations")
)
