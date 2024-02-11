package sqlite

import (
	"database/sql"

	"github.com/dmitrymomot/go-app-template/db"
	_ "github.com/mattn/go-sqlite3" // init sqlite3 driver
)

// Connect establishes a connection to a SQLite database.
// If dbConnString is empty, a connection to an in-memory database is created.
// It returns a pointer to the sql.DB object and an error if any.
func Connect(dbConnString string) (*sql.DB, error) {
	if dbConnString == "" {
		dbConnString = ":memory:"
	}
	return db.InitDB("sqlite3", dbConnString, 1, 1)
}
