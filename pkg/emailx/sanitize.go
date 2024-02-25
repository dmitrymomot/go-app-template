package emailx

import "strings"
import "braces.dev/errtrace"

// SanitizeEmail cleans email address from dots, dashes, etc
func SanitizeEmail(s string, validate bool) (string, error) {
	return errtrace.Wrap2(sanitizeEmail(s, validate))
}

// sanitizeEmail cleans email address from dots, dashes, etc
func sanitizeEmail(s string, validate bool) (string, error) {
	email, err := parseEmail(strings.ToLower(strings.TrimSpace(s)))
	if err != nil {
		return "", errtrace.Wrap(err)
	}
	username := email.LocalPart()
	domain := email.Domain()

	// remove dots, dashes, etc
	{
		if strings.Contains(username, "+") {
			p := strings.Split(username, "+")
			username = p[0]
		}
		if strings.Contains(username, ".") {
			p := strings.Split(username, ".")
			username = strings.Join(p, "")
		}
		if strings.Contains(username, "-") {
			p := strings.Split(username, "-")
			username = strings.Join(p, "")
		}
	}

	newEmail := NewEmailAddress(username, domain)
	if validate {
		if err := ValidateEmail(newEmail.String()); err != nil {
			return newEmail.String(), errtrace.Wrap(err)
		}
	}

	return newEmail.String(), nil
}
