package auth

import "errors"

// Predefined errors
var (
	ErrEmailAlreadyExists              = errors.New("email already exists")
	ErrUserNotFound                    = errors.New("user not found")
	ErrInvalidCredentials              = errors.New("invalid credentials")
	ErrFailedToCreateUser              = errors.New("failed to create user")
	ErrFailedToSendEmail               = errors.New("failed to send email")
	ErrFailedToAuthenticate            = errors.New("failed to authenticate")
	ErrFailedToRestoreAccess           = errors.New("failed to restore access to the account")
	ErrFailedToResetPassword           = errors.New("failed to reset password")
	ErrInvalidToken                    = errors.New("invalid token")
	ErrFailedToSignup                  = errors.New("failed to signup")
	ErrFailedToVerifyEmail             = errors.New("failed to verify email")
	ErrFailedToGetUserProfile          = errors.New("failed to get user profile")
	ErrFailedToReadResponseBody        = errors.New("failed to read response body")
	ErrFailedToAuthenticateUser        = errors.New("failed to authenticate user")
	ErrFailedToCreateUserSocialProfile = errors.New("failed to create user social profile")
	ErrEmptyOrNotVerifiedEmail         = errors.New("email is empty or not verified")
)
