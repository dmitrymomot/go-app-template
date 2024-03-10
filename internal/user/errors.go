package user

import "errors"

// Predefined errors
var (
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrUserNotFound           = errors.New("user not found")
	ErrFailedToRetrieveUser   = errors.New("failed to retrieve user")
	ErrInvalidCredentials     = errors.New("invalid credentials")
	ErrFailedToCreateUser     = errors.New("failed to create user")
	ErrFailedToUpdateEmail    = errors.New("failed to update email")
	ErrFailedToUpdatePassword = errors.New("failed to update password")
	ErrFailedToResetPassword  = errors.New("failed to reset password")
	ErrFailedToVerifyEmail    = errors.New("failed to verify email")
)
