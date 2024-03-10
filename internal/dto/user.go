package dto

import (
	"time"

	"github.com/dmitrymomot/go-app-template/db/repository"
	"github.com/google/uuid"
)

// User represents a user entity.
type User struct {
	ID         uuid.UUID
	Email      string
	CreatedAt  time.Time
	IsVerified bool
}

// CastFromRepositoryUser converts a repository.User object to a User object.
// It takes a repository.User as input and returns a User object along with an error, if any.
func CastFromRepositoryUser(u repository.User) (User, error) {
	uid, err := uuid.Parse(u.ID)
	if err != nil {
		return User{}, err
	}

	return User{
		ID:         uid,
		Email:      u.Email,
		CreatedAt:  u.CreatedAt,
		IsVerified: u.VerifiedAt.Valid && !u.VerifiedAt.Time.IsZero(),
	}, nil
}
