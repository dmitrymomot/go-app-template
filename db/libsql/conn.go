package libsql

import (
	"database/sql"

	"github.com/dmitrymomot/go-app-template/db"
	_ "github.com/tursodatabase/libsql-client-go/libsql" // init libsql driver (SQLite fork)
)

// Connect opens libSQL database connection.
// Can be used also for SQLite3 database connection (libsql driver).
func Connect(dbConnString string, dbMaxOpenConns, dbMaxIdleConns int) (*sql.DB, error) {
	if dbConnString == "" {
		return nil, db.ErrEmptyDBConnString
	}
	return db.InitDB("libsql", dbConnString, dbMaxOpenConns, dbMaxIdleConns)
}
