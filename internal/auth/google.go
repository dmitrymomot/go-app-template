package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/go-app-template/db/repository"
	"github.com/dmitrymomot/go-app-template/internal/dto"
	"github.com/google/uuid"
)

// GoogleService represents a service for handling Google authentication.
type GoogleService struct {
	repo           repository.Querier
	userSvc        googleUserService
	googleClientID string
	providerName   string
}

// Define the user's profile.
type googleProfile struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture,omitempty"`
}

type googleUserService interface {
	CreateUser(ctx context.Context, email, password string) (dto.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (dto.User, error)
	GetUserByEmail(ctx context.Context, email string) (dto.User, error)
}

// NewGoogleService creates a new instance of the GoogleService.
// It takes a repository.Querier and a googleUserService as parameters and returns a pointer to the GoogleService.
func NewGoogleService(repo repository.Querier, userSvc googleUserService, gci string) *GoogleService {
	return &GoogleService{
		repo:           repo,
		userSvc:        userSvc,
		googleClientID: gci,
		providerName:   "google",
	}
}

// Auth is a method to authenticate user by google token.
// It takes a context.Context and a string as parameters and returns a dto.User and an error.
func (s *GoogleService) Auth(ctx context.Context, token string) (dto.User, error) {
	// Get the user's google profile.
	profile, err := s.getUserProfile(token)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToAuthenticate, err))
	}

	// Get the user's social profile by google id and provider type.
	extProfile, err := s.repo.GetUserSocialProfileBySocialID(ctx, repository.GetUserSocialProfileBySocialIDParams{
		ExternalAccountID: profile.ID,
		ProviderType:      s.providerName,
		ProviderID:        s.googleClientID,
	})
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToAuthenticate, err))
		}

		// If the user social profile does not exist, go to the signup process.
		// Create a new user and user social profile.
		u, err := s.signup(ctx, profile)
		if err != nil {
			return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToAuthenticate, err))
		}

		return u, nil
	}

	// Get the user by the user social profile.
	u, err := s.getUserByID(ctx, extProfile.UserID)
	if err != nil {
		return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToAuthenticate, err))
	}

	return u, nil
}

// signup is a method to create a new user by google token.
// It takes a context.Context and a string as parameters and returns a dto.User and an error.
func (s *GoogleService) signup(ctx context.Context, profile googleProfile) (dto.User, error) {
	if profile.Email == "" || !profile.VerifiedEmail {
		return dto.User{}, errtrace.Wrap(ErrEmptyOrNotVerifiedEmail)
	}

	// Check if the user already exists by email.
	// If the user exists, return the user.
	u, err := s.userSvc.GetUserByEmail(ctx, profile.Email)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return dto.User{}, errtrace.Wrap(err)
		}

		// Create a new user.
		u, err = s.userSvc.CreateUser(ctx, profile.Email, "")
		if err != nil {
			return dto.User{}, errtrace.Wrap(errors.Join(ErrFailedToSignup, err))
		}
	}

	// Create a new user social profile.
	if err := s.createSocialProfile(ctx, u.ID, profile); err != nil {
		return dto.User{}, errtrace.Wrap(err)
	}

	return u, nil
}

// getUserProfile retrieves the user's profile from Google using the provided access token.
// It makes an HTTP GET request to the Google API endpoint and parses the response into a googleProfile struct.
// If successful, it returns the user's profile and nil error. Otherwise, it returns an empty googleProfile and an error.
func (s *GoogleService) getUserProfile(token string) (googleProfile, error) {
	// Get the user's profile.
	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return googleProfile{}, errtrace.Wrap(errors.Join(ErrFailedToGetUserProfile, err))
	}
	defer response.Body.Close()

	// Parse the user's profile.
	var profile googleProfile
	if err := json.NewDecoder(response.Body).Decode(&profile); err != nil {
		return googleProfile{}, errtrace.Wrap(errors.Join(ErrFailedToReadResponseBody, err))
	}

	return profile, nil
}

// createSocialProfile creates a new user social profile for the given user ID and Google profile.
// It uses the provided context and the GoogleService's repository to store the social profile information.
// The function returns an error if there was a problem creating the user social profile.
func (s *GoogleService) createSocialProfile(ctx context.Context, u uuid.UUID, profile googleProfile) error {
	// Create a new user social profile.
	if err := s.repo.CreateUserSocialProfile(ctx, repository.CreateUserSocialProfileParams{
		UserID:            u.String(),
		ProviderID:        s.googleClientID,
		ProviderType:      s.providerName,
		ExternalAccountID: profile.ID,
	}); err != nil {
		return errtrace.Wrap(errors.Join(ErrFailedToCreateUserSocialProfile, err))
	}

	return nil
}

// getUserByID retrieves a user by their ID.
// It takes a context and an ID as input parameters and returns a dto.User and an error.
func (s *GoogleService) getUserByID(ctx context.Context, id string) (dto.User, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return dto.User{}, errtrace.Wrap(err)
	}
	u, err := s.userSvc.GetUserByID(ctx, uid)
	if err != nil {
		return dto.User{}, errtrace.Wrap(err)
	}

	return u, nil
}
