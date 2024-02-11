package migration

import "errors"

// Predefined errors.
var (
	ErrFailedToApplyMigrations = errors.New("failed to apply migrations")
	ErrMissedDBConnection      = errors.New("missed db connection")
	ErrUndefinedDBDriver       = errors.New("undefined db driver")
)
