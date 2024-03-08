package token

import (
	"errors"
	"time"

	"github.com/dmitrymomot/go-app-template/internal/dto"
	"github.com/dmitrymomot/go-signature"
)

// VerificationService provides functionality for verifying tokens.
type VerificationService struct {
	secretKey string
}

// NewVerificationService creates a new instance of the VerificationService.
// It takes a secretKey as a parameter and returns a pointer to the VerificationService.
func NewVerificationService(secretKey string) *VerificationService {
	return &VerificationService{
		secretKey: secretKey,
	}
}

// Generate generates a new token with the specified payload and expiration time.
// It takes a payload and expiration time as parameters and returns a token and an error.
func (s *VerificationService) Generate(payload dto.VerificationToken, expiration time.Duration) (string, error) {
	signature.SetSigningKey(s.secretKey)
	token, err := signature.NewTemporary(payload, expiration)
	if err != nil {
		return "", errors.Join(ErrFailedToGenerateToken, err)
	}
	return token, nil
}

// Verify verifies the specified token.
// It takes a token as a parameter and returns a payload and an error.
func (s *VerificationService) Verify(token string) (dto.VerificationToken, error) {
	signature.SetSigningKey(s.secretKey)
	payload, err := signature.Parse[dto.VerificationToken](token)
	if err != nil {
		return payload, errors.Join(ErrFailedToVerifyToken, err)
	}
	return payload, nil
}
