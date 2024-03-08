package user

import "errors"

// Predefined errors
var (
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrUserNotFound           = errors.New("user not found")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrFailedToCreateUser     = errors.New("failed to create user")
	ErrFailedToUpdateEmail    = errors.New("failed to update email")
	ErrFailedToUpdatePassword = errors.New("failed to update password")
)
