package main

import (
	"flag"

	"github.com/dmitrymomot/go-app-template/db/libsql"
	"github.com/dmitrymomot/go-app-template/db/migration"
)

func main() {
	// Parse flags
	rollback := flag.Bool("rollback", false, "Rollback all migrations")
	flag.Parse()

	log := initLogger()
	defer log.Sync() //nolint:errcheck
	logger := log.Sugar()
	logger.Info("Starting db migration...")

	// Init db connection
	db, err := libsql.Connect(dbConnString, 1, 1)
	if err != nil {
		logger.Fatalw("Failed to open db connection", "error", err)
	}
	defer db.Close()

	// Rollback all migrations
	if rollback != nil && *rollback {
		n, err := migration.Down(db, "sqlite3", migrationsTable, migrationsDir)
		if err != nil {
			logger.Fatalw("Failed to rollback migrations", "error", err)
		}
		logger.Infof("Rolled back %d migrations!", n)
		return
	}

	// Apply all migrations
	n, err := migration.Up(db, "sqlite3", migrationsTable, migrationsDir)
	if err != nil {
		logger.Fatalw("Failed to apply migrations", "error", err)
	}
	logger.Infof("Applied %d migrations!", n)
}
