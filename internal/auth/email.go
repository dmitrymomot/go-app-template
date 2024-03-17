package auth

import (
	"context"
	"errors"
	"time"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/go-app-template/internal/dto"
	"github.com/google/uuid"
)

// Service represents an authentication service.
type EmailService struct {
	userSvc  userService
	mailSvc  emailSender
	tokenSvc tokenService

	resetPasswordTTL time.Duration
	confirmEmailTTL  time.Duration
}

// userService is an interface for user operations.
type userService interface {
	CreateUser(ctx context.Context, email, password string) (dto.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (dto.User, error)
	GetUserByEmail(ctx context.Context, email string) (dto.User, error)
	CheckPasswordByEmail(ctx context.Context, email, password string) (dto.User, error)
	ResetUserPassword(ctx context.Context, id uuid.UUID, password string) error
	SetVerificationStatus(ctx context.Context, id uuid.UUID, verified bool) error
}

// emailSender is an interface for sending emails.
type emailSender interface {
	SendWelcomeEmail(ctx context.Context, userID, email, token string) error
	SendPasswordResetEmail(ctx context.Context, userID, email, token string) error
	SendVerificationEmail(ctx context.Context, userID, email, token string) error
}

// tokenService is an interface for token operations.
type tokenService interface {
	Generate(payload dto.Verification, expiration time.Duration) (string, error)
	Verify(token string) (dto.Verification, error)
}

// NewEmailService creates a new instance of the auth EmailService struct.
// It takes a repository.Querier as a parameter and returns a pointer to the Service.
func NewEmailService(
	userSvc userService,
	mailSvc emailSender,
	tokenSvc tokenService,
	resetPasswordTTL time.Duration,
	confirmEmailTTL time.Duration,
) *EmailService {
	if resetPasswordTTL == 0 {
		resetPasswordTTL = time.Minute * 15
	}
	if confirmEmailTTL == 0 {
		confirmEmailTTL = time.Hour * 24
	}
	return &EmailService{
		userSvc:          userSvc,
		mailSvc:          mailSvc,
		tokenSvc:         tokenSvc,
		resetPasswordTTL: resetPasswordTTL,
		confirmEmailTTL:  confirmEmailTTL,
	}
}

// Signup is a method that creates a new user account.
// It takes an email and a password as parameters and returns an error.
func (s *EmailService) Signup(ctx context.Context, email, password string) error {
	user, err := s.userSvc.CreateUser(ctx, email, password)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToSignup, err))
	}

	token, err := s.tokenSvc.Generate(dto.NewEmailVerification(
		user.ID.String(),
		user.Email,
	), time.Hour*24)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToSignup, err))
	}

	if err := s.mailSvc.SendWelcomeEmail(ctx, user.ID.String(), user.Email, token); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToSignup, err))
	}

	return nil
}

// Login is a method that authenticates a user.
// It takes an email and a password as parameters and returns an error.
func (s *EmailService) Login(ctx context.Context, email, password string) (dto.User, error) {
	user, err := s.userSvc.CheckPasswordByEmail(ctx, email, password)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToAuthenticate, err))
	}

	return user, nil
}

// ForgotPassword is a method that sends a password reset email.
// It takes an email as a parameter and returns an error.
func (s *EmailService) ForgotPassword(ctx context.Context, email string) error {
	user, err := s.userSvc.GetUserByEmail(ctx, email)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrUserNotFound, err))
	}

	token, err := s.tokenSvc.Generate(dto.NewPasswordResetTokenVerification(
		user.ID.String(),
		user.Email,
	), time.Minute*15)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToRestoreAccess, err))
	}

	if err := s.mailSvc.SendPasswordResetEmail(ctx, user.ID.String(), email, token); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToRestoreAccess, err))
	}

	return nil
}

// ResetPassword is a method that resets a user's password.
// It takes a token and a new password as parameters and returns an error.
func (s *EmailService) ResetPassword(ctx context.Context, token, password string) error {
	payload, err := s.tokenSvc.Verify(token)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToResetPassword, err))
	}

	if payload.Type != dto.PasswordResetVerificationType {
		return errtrace.Wrap(errors.Join(ErrFailedToResetPassword, ErrInvalidToken))
	}

	uid, err := uuid.Parse(payload.UserID)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToResetPassword, err))
	}
	if err := s.userSvc.ResetUserPassword(ctx, uid, password); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToResetPassword, err))
	}

	return nil
}

// VerifyEmail is a method that verifies a user's email.
// It takes a token as a parameter and returns an error.
func (s *EmailService) VerifyEmail(ctx context.Context, token string) (dto.User, error) {
	payload, err := s.tokenSvc.Verify(token)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToVerifyEmail, err))
	}

	if payload.Type != dto.EmailVerificationType {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToVerifyEmail, ErrInvalidToken))
	}

	user, err := s.userSvc.GetUserByEmail(ctx, payload.Email)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToVerifyEmail, err))
	}

	if user.ID.String() != payload.UserID {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToVerifyEmail, ErrInvalidToken))
	}

	if err := s.userSvc.SetVerificationStatus(ctx, user.ID, true); err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToVerifyEmail, err))
	}

	return user, nil
}
