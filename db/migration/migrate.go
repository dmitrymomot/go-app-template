package migration

import (
	"database/sql"
	"errors"

	migrate "github.com/rubenv/sql-migrate"
)

// Up applies all available migrations.
func Up(db *sql.DB, driver, migrationsTable, migrationsDir string) (int, error) {
	return Run(db, driver, migrationsTable, migrationsDir, migrate.Up)
}

// Down rolls back all migrations.
func Down(db *sql.DB, driver, migrationsTable, migrationsDir string) (int, error) {
	return Run(db, driver, migrationsTable, migrationsDir, migrate.Down)
}

// Run executes database migrations using the provided parameters.
// It takes a *sql.DB object, the database driver name, the name of the migrations table,
// the directory containing the migration files, and the migration direction.
// It returns the number of applied migrations and any error encountered.
func Run(db *sql.DB, driver, migrationsTable, migrationsDir string, direction migrate.MigrationDirection) (int, error) {
	// Validate input parameters
	if db == nil {
		return 0, ErrMissedDBConnection
	}
	if driver == "" {
		return 0, ErrUndefinedDBDriver
	}
	if migrationsTable == "" {
		migrationsTable = "migrations"
	}
	if migrationsDir == "" {
		migrationsTable = "./db/sql/migrations"
	}

	// Create migration set
	m := migrate.MigrationSet{
		TableName: migrationsTable,
	}
	migrations := &migrate.FileMigrationSource{
		Dir: migrationsDir,
	}

	// Apply migrations
	n, err := m.Exec(db, driver, migrations, direction)
	if err != nil {
		return 0, errors.Join(ErrFailedToApplyMigrations, err)
	}

	return n, nil
}
