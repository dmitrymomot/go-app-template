package auth

import "github.com/dmitrymomot/go-app-template/db/repository"

// Service represents an authentication service.
type EmailService struct {
	repo repository.Querier
}

// NewEmailService creates a new instance of the auth EmailService struct.
// It takes a repository.Querier as a parameter and returns a pointer to the Service.
func NewEmailService(repo repository.Querier) *EmailService {
	return &EmailService{
		repo: repo,
	}
}
