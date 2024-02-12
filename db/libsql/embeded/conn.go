package libsql_embeded

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"braces.dev/errtrace"
	"github.com/libsql/go-libsql" // init libSQL driver
)

// Connect establishes a connection to a libSQL/SQLite database.
// If dbConnString is empty, a connection to an in-memory database is created.
// It returns a pointer to the sql.DB object and an error if any.
func Connect(dbName, primaryUrl, authToken, tempDir string) (*sql.DB, error) {
	if dbName == "" {
		return nil, errtrace.Wrap(ErrMissedDBName)
	}
	if primaryUrl == "" {
		return nil, errtrace.Wrap(ErrMissedPrimaryURL)
	}
	if authToken == "" {
		return nil, errtrace.Wrap(ErrMissedAuthToken)
	}

	// Create a temporary directory to store the database file
	dir, err := os.MkdirTemp(tempDir, "libsql-*")
	if err != nil {
		return nil, errtrace.Wrap(errors.Join(ErrFailedToCreateTempDir, err))
	}
	defer os.RemoveAll(dir) //nolint:errcheck

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, authToken)
	if err != nil {
		return nil, errtrace.Wrap(errors.Join(ErrFailedToCreateConnector, err))
	}
	defer connector.Close()

	return sql.OpenDB(connector), nil
}
