package auth

import "github.com/dmitrymomot/go-app-template/db/repository"

// Service represents an authentication service.
type Service struct {
	repo repository.Querier
}

// NewService creates a new instance of the auth Service struct.
// It takes a repository.Querier as a parameter and returns a pointer to the Service.
func NewService(repo repository.Querier) *Service {
	return &Service{
		repo: repo,
	}
}
