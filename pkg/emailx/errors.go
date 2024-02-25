package emailx

import "errors"

// Predefined errors.
var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrInvalidIcanSuffix  = errors.New("invalid ICAN suffix")
	ErrInvalidEmailHost   = errors.New("invalid email host")
)
