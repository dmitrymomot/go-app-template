package main

import (
	"database/sql"
	"flag"

	migrate "github.com/rubenv/sql-migrate"
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
	db, err := sql.Open("libsql", dbConnString)
	if err != nil {
		logger.Fatalw("Failed to open db connection", "error", err)
	}
	defer db.Close()

	// check db connection
	if err := db.Ping(); err != nil {
		logger.Fatalw("Failed to ping db", "error", err)
	}

	m := migrate.MigrationSet{
		TableName: migrationsTable,
	}
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	// Apply migrations
	if rollback != nil && *rollback {
		n, err := m.Exec(db, "sqlite3", migrations, migrate.Down)
		if err != nil {
			logger.Fatalw("Failed to rollback migrations", "error", err)
		}
		logger.Infof("Rolled back %d migrations!", n)
		return
	}

	n, err := m.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		logger.Fatalw("Failed to apply migrations", "error", err)
	}
	logger.Infof("Applied %d migrations!", n)
}
