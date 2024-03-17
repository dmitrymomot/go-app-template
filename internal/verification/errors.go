package verification

import "errors"

// Predefined token types
var (
	ErrFailedToGenerateToken = errors.New("failed to generate token")
	ErrFailedToVerifyToken   = errors.New("failed to verify token")
)
