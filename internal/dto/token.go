package dto

// Predefined token types
const (
	EmailVerificationTokenType      = "email_verification"
	PasswordResetTokenType          = "password_reset"
	DeleteUserConfirmationTokenType = "delete_user_confirmation"
)

// VerificationToken represents a verification token entity.
type VerificationToken struct {
	Type   string `json:"t"`
	UserID string `json:"u"`
	Email  string `json:"e"`
}

// NewEmailVerificationToken creates a new instance of the VerificationToken struct for email verification.
// It takes a user ID and an email as parameters and returns a pointer to the VerificationToken.
func NewEmailVerificationToken(userID, email string) VerificationToken {
	return VerificationToken{
		Type:   EmailVerificationTokenType,
		UserID: userID,
		Email:  email,
	}
}

// NewPasswordResetToken creates a new instance of the VerificationToken struct for password reset.
// It takes a user ID and an email as parameters and returns a pointer to the VerificationToken.
func NewPasswordResetToken(userID, email string) VerificationToken {
	return VerificationToken{
		Type:   PasswordResetTokenType,
		UserID: userID,
		Email:  email,
	}
}

// NewDeleteUserConfirmationToken creates a new instance of the VerificationToken struct for user deletion confirmation.
// It takes a user ID and an email as parameters and returns a pointer to the VerificationToken.
func NewDeleteUserConfirmationToken(userID, email string) VerificationToken {
	return VerificationToken{
		Type:   DeleteUserConfirmationTokenType,
		UserID: userID,
		Email:  email,
	}
}
