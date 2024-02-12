package libsql_local

import (
	"database/sql"

	"github.com/dmitrymomot/go-app-template/db"
	_ "github.com/libsql/go-libsql" // init libSQL driver (it's fully compatible with the sqlite3)
)

// Connect establishes a connection to a libSQL/SQLite database.
// If dbConnString is empty, a connection to an in-memory database is created.
// It returns a pointer to the sql.DB object and an error if any.
func Connect(dbConnString string) (*sql.DB, error) {
	if dbConnString == "" {
		dbConnString = ":memory:"
	}
	return db.InitDB("libsql", dbConnString, 1, 1)
}
