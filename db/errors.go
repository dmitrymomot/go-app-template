package db

import "errors"

// Predefined errors.
var (
	ErrFailedToOpenDBConnection = errors.New("failed to open db connection")
	ErrFailedToPingDB           = errors.New("failed to ping db")
	ErrEmptyDBConnString        = errors.New("empty db connection string")
	ErrUndefinedDBDriver        = errors.New("undefined db driver")
)
