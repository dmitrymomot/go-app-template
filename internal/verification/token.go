package verification

import (
	"errors"
	"time"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/go-app-template/internal/dto"
	"github.com/dmitrymomot/go-signature/v2"
)

// TokenVerificationService provides functionality for verifying tokens.
type TokenVerificationService struct {
	signer tokenSigner[dto.Verification]
}

// tokenSigner is an interface that defines the methods for signing and parsing data.
type tokenSigner[Payload any] interface {
	// SignTemporary generates a temporary signature for the given data with a specified time-to-live (TTL).
	// It returns the generated signature as a string and any error encountered.
	SignTemporary(data Payload, ttl time.Duration) (string, error)

	// Parse parses the given token and returns the payload associated with it.
	// It returns the parsed payload and any error encountered.
	Parse(token string) (Payload, error)
}

// NewTokenVerificationService creates a new instance of the VerificationService.
// It takes a secretKey as a parameter and returns a pointer to the VerificationService.
func NewTokenVerificationService(secretKey []byte) *TokenVerificationService {
	return &TokenVerificationService{
		signer: signature.NewSigner256[dto.Verification](secretKey),
	}
}

// Generate generates a new token with the specified payload and expiration time.
// It takes a payload and expiration time as parameters and returns a token and an error.
func (s *TokenVerificationService) Generate(payload dto.Verification, expiration time.Duration) (string, error) {
	token, err := s.signer.SignTemporary(payload, expiration)
	if err != nil {
		return "", errtrace.Wrap(errors.Join(ErrFailedToGenerateToken, err))
	}
	return token, nil
}

// Verify verifies the specified token.
// It takes a token as a parameter and returns a payload and an error.
func (s *TokenVerificationService) Verify(token string) (dto.Verification, error) {
	payload, err := s.signer.Parse(token)
	if err != nil {
		return payload, errtrace.Wrap(errors.Join(ErrFailedToVerifyToken, err))
	}
	return payload, nil
}
