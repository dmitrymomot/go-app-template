package user

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/go-app-template/db/repository"
	"github.com/dmitrymomot/go-app-template/internal/dto"
	"github.com/dmitrymomot/go-app-template/pkg/emailx"
	"github.com/dmitrymomot/go-app-template/pkg/validator"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Service represents a user service.
type Service struct {
	repo repository.Querier
}

// NewService creates a new instance of the user Service struct.
// It takes a repository.Querier as a parameter and returns a pointer to the Service.
func NewService(repo repository.Querier) *Service {
	return &Service{
		repo: repo,
	}
}

// CreateUser creates a new user with the specified email and password.
// It sanitizes the email, checks if the email is already in use, prepares the user data,
// and creates a new user in the repository.
// The function returns the UUID of the created user and any error encountered during the process.
func (s *Service) CreateUser(ctx context.Context, email, password string) (uuid.UUID, error) {
	// Sanitize the email
	email, err := emailx.SanitizeEmail(email, true)
	if err != nil {
		return uuid.Nil, errtrace.Wrap(errors.Join(ErrFailedToCreateUser, err))
	}

	// Check if the email is already in use
	if _, err := s.repo.GetUserByEmail(ctx, email); err == nil {
		return uuid.Nil, errtrace.Wrap(errors.Join(ErrFailedToCreateUser, ErrEmailAlreadyExists))
	}

	// Prepare a new user data
	uid := uuid.New()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return uuid.Nil, errtrace.Wrap(errors.Join(ErrFailedToCreateUser, err))
	}

	// Create a new user
	if err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		ID:        uid.String(),
		Email:     email,
		Password:  sql.NullString{String: string(passwordHash), Valid: true},
		CreatedAt: time.Now(),
	}); err != nil {
		return uuid.Nil, errtrace.Wrap(errors.Join(ErrFailedToCreateUser, err))
	}

	return uid, nil
}

// GetUserByEmail returns a user with the specified email.
// The function returns the user data and any error encountered during the process.
func (s *Service) GetUserByEmail(ctx context.Context, email string) (dto.User, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrUserNotFound, err))
	}

	uid, err := uuid.Parse(user.ID)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrUserNotFound, err))
	}

	return dto.User{
		ID:        uid,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

// GetUserByID returns a user with the specified ID.
// The function returns the user data and any error encountered during the process.
func (s *Service) GetUserByID(ctx context.Context, id uuid.UUID) (dto.User, error) {
	// Get the user by ID
	user, err := s.repo.GetUserByID(ctx, id.String())
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrUserNotFound, err))
	}

	// Parse the user ID
	uid, err := uuid.Parse(user.ID)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrUserNotFound, err))
	}

	return dto.User{
		ID:        uid,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}, nil
}

// UpdateUserEmail updates a user email with the specified ID.
// The function returns any error encountered during the process.
func (s *Service) UpdateUserEmail(ctx context.Context, id uuid.UUID, email string) error {
	// Sanitize the email
	email, err := emailx.SanitizeEmail(email, true)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToUpdateEmail, err))
	}

	// Update the user
	if err := s.repo.UpdateUserEmail(ctx, repository.UpdateUserEmailParams{
		ID:    id.String(),
		Email: email,
	}); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToUpdateEmail, err))
	}

	return nil
}

// UpdateUserPassword updates a user password with the specified ID.
// The function returns any error encountered during the process.
func (s *Service) UpdateUserPassword(ctx context.Context, id uuid.UUID, currentPassword, newPassword string) error {
	// Check the current password
	if err := s.CheckPasswordByID(ctx, id, currentPassword); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToUpdatePassword, err))
	}

	// Validate the new password
	if err := validator.ValidatePassword(newPassword); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToUpdatePassword, err))
	}

	// Prepare a new password hash
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToUpdatePassword, err))
	}

	// Update the user
	if err := s.repo.UpdateUserPassword(ctx, repository.UpdateUserPasswordParams{
		ID:       id.String(),
		Password: sql.NullString{String: string(passwordHash), Valid: true},
	}); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToUpdatePassword, err))
	}

	return nil
}

// DeleteUser deletes a user with the specified ID.
// The function returns any error encountered during the process.
func (s *Service) DeleteUser(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteUser(ctx, id.String()); err != nil {
		return errtrace.Wrap(errors.Join(ErrUserNotFound, err))
	}

	return nil
}

// CheckPasswordByID checks if the provided password matches the user's password.
// The function returns any error encountered during the process.
func (s *Service) CheckPasswordByID(ctx context.Context, id uuid.UUID, password string) error {
	user, err := s.repo.GetUserByID(ctx, id.String())
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrInvalidCredentials, err))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password)); err != nil {
		return errtrace.Wrap(errors.Join(ErrInvalidCredentials, err))
	}

	return nil
}

// CheckPasswordByEmail checks if the provided password matches the user's password.
// The function returns any error encountered during the process.
func (s *Service) CheckPasswordByEmail(ctx context.Context, email, password string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return errtrace.Wrap(errors.Join(ErrInvalidCredentials, err))
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(password)); err != nil {
		return errtrace.Wrap(errors.Join(ErrInvalidCredentials, err))
	}

	return nil
}
