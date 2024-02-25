package emailx

import (
	"fmt"
	"net"
	"strings"

	"braces.dev/errtrace"
	"golang.org/x/net/publicsuffix"
)

// ValidateEmail validates the email address locally. If you want to validate the host of
// this email address remotely call the ValidateHost method.
func ValidateEmail(emailAddr string) error {
	// Parse the email address and validate the format
	email, err := parseEmail(emailAddr)
	if err != nil {
		return errtrace.Wrap(err)
	}

	// Check if the email host is valid
	if err := ValidateHost(email.Domain()); err != nil {
		return errtrace.Wrap(err)
	}

	// Check if the email domain is managed by ICANN
	if err := ValidateIcanSuffix(email.Domain()); err != nil {
		return errtrace.Wrap(err)
	}

	return nil
}

// ValidateHost will test if the email address is actually reachable. It will first try to resolve
// the host and then start a mail transaction.
func ValidateHost(host string) error {
	if _, err := lookupHost(host); err != nil {
		return errtrace.Wrap(fmt.Errorf("%w: %s", ErrInvalidEmailHost, err))
	}
	return nil
}

// ValidateIcanSuffix will test if the public suffix of the domain is managed by ICANN using
// the golang.org/x/net/publicsuffix package. If not it will return an error. Note that if this
// method returns an error it does not necessarily mean that the email address is invalid. Also the
// suffix list in the standard package is embedded and thereby not up to date.
func ValidateIcanSuffix(host string) error {
	d := strings.ToLower(host)
	if _, icann := publicsuffix.PublicSuffix(d); !icann {
		return errtrace.Wrap(ErrInvalidIcanSuffix)
	}
	return nil
}

// lookupHost first checks if any MX records are available and if not, it will check
// if A records are available as they can resolve email server hosts. An error indicates
// that non of the A or MX records are available.
func lookupHost(domain string) (string, error) {
	if mx, err := net.LookupMX(domain); err == nil {
		return mx[0].Host, nil
	}
	if ips, err := net.LookupIP(domain); err == nil {
		return ips[0].String(), nil // randomly returns IPv4 or IPv6 (when available)
	}
	return "", errtrace.Wrap(fmt.Errorf("failed finding MX and A records for domain %s", domain))
}
