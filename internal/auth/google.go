package auth

import "github.com/dmitrymomot/go-app-template/db/repository"

type GoogleService struct {
	repo repository.Querier
}

func NewGoogleService(repo repository.Querier) *GoogleService {
	return &GoogleService{
		repo: repo,
	}
}
