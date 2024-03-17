package dto

// Predefined token types
const (
	// EmailVerificationType is the type for email verification tokens.
	EmailVerificationType = "email_verification"
	// PasswordResetVerificationType is the type for password reset tokens.
	PasswordResetVerificationType = "password_reset"
	// DeleteUserVerificationType is the type for user deletion confirmation tokens.
	DeleteUserVerificationType = "delete_user_confirmation"
)

// Verification represents a verification token entity.
type Verification struct {
	Type   string `json:"t"`
	UserID string `json:"u"`
	Email  string `json:"e"`
}

// NewEmailVerification creates a new instance of the Verification struct for email verification.
// It takes a user ID and an email as parameters and returns a pointer to the Verification.
func NewEmailVerification(userID, email string) Verification {
	return Verification{
		Type:   EmailVerificationType,
		UserID: userID,
		Email:  email,
	}
}

// NewPasswordResetTokenVerification creates a new instance of the Verification struct for password reset.
// It takes a user ID and an email as parameters and returns a pointer to the Verification.
func NewPasswordResetTokenVerification(userID, email string) Verification {
	return Verification{
		Type:   PasswordResetVerificationType,
		UserID: userID,
		Email:  email,
	}
}

// NewDeleteUserTokenVerification creates a new instance of the Verification struct for user deletion confirmation.
// It takes a user ID and an email as parameters and returns a pointer to the Verification.
func NewDeleteUserTokenVerification(userID, email string) Verification {
	return Verification{
		Type:   DeleteUserVerificationType,
		UserID: userID,
		Email:  email,
	}
}
