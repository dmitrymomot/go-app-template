package main

import (
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

var (
	// App
	appName      = env.GetString("APP_NAME", "go-app-template")
	appEnv       = env.GetString("APP_ENV", "local") // local, development, production, testing, staging
	appDebugMode = env.GetBool("APP_DEBUG_MODE", false)
	appLogLevel  = env.GetString("APP_LOG_LEVEL", "info") // debug, info, warn, error

	// Build
	buildTag = env.GetString("COMMIT_HASH", "undefined")

	// DB
	dbConnString    = env.MustString("DATABASE_URL")
	migrationsDir   = env.GetString("DATABASE_MIGRATIONS_DIR", "./db/sql/migrations")
	migrationsTable = env.GetString("DATABASE_MIGRATIONS_TABLE", "migrations")
)
