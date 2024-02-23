package libsql_remote

import (
	"database/sql"

	"braces.dev/errtrace" // init libsql driver (SQLite fork)
	"github.com/dmitrymomot/go-app-template/db"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

// Connect opens libSQL database connection.
// Can be used also for SQLite3 database connection (libsql driver).
func Connect(dbConnString string, dbMaxOpenConns, dbMaxIdleConns int) (*sql.DB, error) {
	if dbConnString == "" {
		return nil, errtrace.Wrap(db.ErrEmptyDBConnString)
	}
	return errtrace.Wrap2(db.InitDB("libsql", dbConnString, dbMaxOpenConns, dbMaxIdleConns))
}
