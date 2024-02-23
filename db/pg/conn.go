package pg

import (
	"database/sql"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/go-app-template/db"
	_ "github.com/lib/pq" // init postgres driver
)

// Connect establishes a connection to the PostgreSQL database using the provided connection string.
// It returns a pointer to the sql.DB object and an error if the connection fails.
// The dbConnString parameter is the connection string for the PostgreSQL database.
func Connect(dbConnString string, dbMaxOpenConns, dbMaxIdleConns int) (*sql.DB, error) {
	return errtrace.Wrap2(db.InitDB("postgres", dbConnString, dbMaxOpenConns, dbMaxIdleConns))
}
