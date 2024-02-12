package libsql_embeded

import (
	"database/sql"
	"errors"
	"os"
	"path/filepath"

	"github.com/libsql/go-libsql" // init libSQL driver
)

// Connect establishes a connection to a libSQL/SQLite database.
// If dbConnString is empty, a connection to an in-memory database is created.
// It returns a pointer to the sql.DB object and an error if any.
func Connect(dbName, primaryUrl, authToken, tempDir string) (*sql.DB, error) {
	if dbName == "" {
		return nil, ErrMissedDBName
	}
	if primaryUrl == "" {
		return nil, ErrMissedPrimaryURL
	}
	if authToken == "" {
		return nil, ErrMissedAuthToken
	}

	// Create a temporary directory to store the database file
	dir, err := os.MkdirTemp(tempDir, "libsql-*")
	if err != nil {
		return nil, errors.Join(ErrFailedToCreateTempDir, err)
	}
	defer os.RemoveAll(dir) //nolint:errcheck

	dbPath := filepath.Join(dir, dbName)

	connector, err := libsql.NewEmbeddedReplicaConnector(dbPath, primaryUrl, authToken)
	if err != nil {
		return nil, errors.Join(ErrFailedToCreateConnector, err)
	}
	defer connector.Close()

	return sql.OpenDB(connector), nil
}
