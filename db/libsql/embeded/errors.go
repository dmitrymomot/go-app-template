package libsql_embeded

import (
	"errors"
)

// Predefined errors
var (
	ErrMissedDBName            = errors.New("empty database name")
	ErrMissedPrimaryURL        = errors.New("empty primary URL")
	ErrMissedAuthToken         = errors.New("empty auth token")
	ErrFailedToCreateTempDir   = errors.New("failed to create a temporary directory")
	ErrFailedToCreateConnector = errors.New("failed to create a connector")
)
