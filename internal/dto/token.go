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
