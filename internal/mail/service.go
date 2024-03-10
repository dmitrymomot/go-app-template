package mail

import (
	"context"

	"github.com/dmitrymomot/mailer"
)

// Service represents a mail service.
type Service struct {
	sender emailSender
}

// emailSender is an interface for sending emails.
type emailSender interface {
	SendEmail(ctx context.Context, payload mailer.SendEmailPayload) error
}

// NewService creates a new instance of the mail Service struct.
// It takes an emailSender as a parameter and returns a pointer to the Service.
func NewService(sender emailSender) *Service {
	return &Service{
		sender: sender,
	}
}

// SendWelcomeEmail is a method that sends a welcome email to a user.
// It takes an email and a name as parameters and returns an error.
func (s *Service) SendWelcomeEmail(ctx context.Context, userID, email, token string) error {
	return nil
}

// SendPasswordResetEmail is a method that sends a password reset email to a user.
// It takes an email and a reset URL as parameters and returns an error.
func (s *Service) SendPasswordResetEmail(ctx context.Context, userID, email, token string) error {
	return nil
}

// SendVerificationEmail is a method that sends a verification email to a user.
// It takes an email and a verification URL as parameters and returns an error.
func (s *Service) SendVerificationEmail(ctx context.Context, userID, email, token string) error {
	return nil
}

// SendUserDestroyEmail is a method that sends a confirmation email to a user.
// It takes an email as a parameter and returns an error.
func (s *Service) SendUserDestroyEmail(ctx context.Context, userID, email, token string) error {
	return nil
}
