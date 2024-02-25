package emailx

import (
	"braces.dev/errtrace"
	"regexp"
	"strings"
)

// ParseEmail parses an email address string.
func ParseEmail(s string) (*EmailAddress, error) {
	return errtrace.Wrap2(parseEmail(s))
}

// regular expression for email validation
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,10}$`)

// Parse will parse the input and validate the email locally. If you want to validate the host of
// this email address remotely call the ValidateHost method.
func parseEmail(email string) (*EmailAddress, error) {
	if !emailRegex.MatchString(email) {
		return nil, errtrace.Wrap(ErrInvalidEmailFormat)
	}

	i := strings.LastIndexByte(email, '@')
	e := NewEmailAddress(email[:i], email[i+1:])

	return e, nil
}
