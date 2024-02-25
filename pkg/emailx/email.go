package emailx

import "fmt"

// EmailAddress is a structure that stores the address local-part@domain parts.
type EmailAddress struct {
	// LocalPart usually the username of an email address.
	localPart string

	// Domain is the part of the email address after the last @.
	// This should be DNS resolvable to an email server.
	domain string
}

// NewEmailAddress creates a new email address.
func NewEmailAddress(localPart, domain string) *EmailAddress {
	return &EmailAddress{
		localPart: localPart,
		domain:    domain,
	}
}

// String returns the email address as a string.
func (e EmailAddress) String() string {
	if e.localPart == "" || e.domain == "" {
		return ""
	}
	return fmt.Sprintf("%s@%s", e.localPart, e.domain)
}

// LocalPart returns the local part of the email address.
func (e EmailAddress) LocalPart() string {
	return e.localPart
}

// Domain returns the domain part of the email address.
func (e EmailAddress) Domain() string {
	return e.domain
}
