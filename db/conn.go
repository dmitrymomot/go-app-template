package db

import (
	"database/sql"
	"errors"

	"braces.dev/errtrace"
)

// InitDB initializes a database connection using the specified driver, connection string,
// maximum open connections, and maximum idle connections.
// It returns a pointer to the sql.DB object and an error if any occurred during the initialization process.
func InitDB(driver, dbConnString string, dbMaxOpenConns, dbMaxIdleConns int) (*sql.DB, error) {
	// Validate input parameters
	if dbConnString == "" {
		return nil, errtrace.Wrap(ErrEmptyDBConnString)
	}
	if driver == "" {
		return nil, errtrace.Wrap(ErrUndefinedDBDriver)
	}
	if dbMaxOpenConns <= 0 {
		dbMaxOpenConns = 1
	}
	if dbMaxIdleConns <= 0 {
		dbMaxIdleConns = 1
	}

	// Init db connection
	db, err := sql.Open(driver, dbConnString)
	if err != nil {
		return nil, errtrace.Wrap(errors.Join(ErrFailedToOpenDBConnection, err))
	}

	// Set db connection pool settings
	if dbMaxOpenConns < dbMaxIdleConns {
		dbMaxOpenConns = dbMaxIdleConns
	}
	if dbMaxIdleConns > 0 {
		db.SetMaxOpenConns(dbMaxOpenConns)
	}
	if dbMaxIdleConns > 0 {
		db.SetMaxIdleConns(dbMaxIdleConns)
	}

	// check db connection
	if err := db.Ping(); err != nil {
		return nil, errtrace.Wrap(errors.Join(ErrFailedToPingDB, err))
	}

	return db, nil
}
