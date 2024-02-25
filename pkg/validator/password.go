package validator

import (
	"braces.dev/errtrace"
	"errors"
	"regexp"
)

// Predefined validation errors.
var (
	ErrPasswordTooShort = errors.New("password is too short, must be at least 8 characters long")
	ErrPasswordTooLong  = errors.New("password is too long, must be no more than 64 characters long")
	ErrPasswordInvalid  = errors.New("password must contain at least one digit, uppercase, and lowercase letters")
)

// regular expression for password validation
var (
	regexpPasswordDigit    = regexp.MustCompile(`[0-9]+`)
	regexpPasswordUppercas = regexp.MustCompile(`[A-Z]+`)
	regexpPasswordLowercas = regexp.MustCompile(`[a-z]+`)
)

// ValidatePassword validates the password.
// password must contain at least 8 characters
// and at least one digit, one uppercase letter, and one lowercase letter.
func ValidatePassword(password string) error {
	// Check if password is too short or too long.
	// This is not necessary because the password length is already validated
	// by the minLen and maxLen tags in the validator package.
	// if len(password) < 8 {
	// 	return ErrPasswordTooShort
	// } else if len(password) > 64 {
	// 	return ErrPasswordTooLong
	// }

	// Check if password contains at least one digit.
	if !regexpPasswordDigit.MatchString(password) {
		return errtrace.Wrap(ErrPasswordInvalid)
	}
	// Check if password contains at least one uppercase letter.
	if !regexpPasswordUppercas.MatchString(password) {
		return errtrace.Wrap(ErrPasswordInvalid)
	}
	// Check if password contains at least one lowercase letter.
	if !regexpPasswordLowercas.MatchString(password) {
		return errtrace.Wrap(ErrPasswordInvalid)
	}

	return nil
}
