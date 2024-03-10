package stringx

import (
	"errors"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// ToSlug converts any string into a slug.
func ToSlug(s string) (string, error) {
	// Normalize string (remove accents) and make lowercase.
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	normalized, _, err := transform.String(t, s)
	if err != nil {
		return "", errors.Join(ErrFailedToNormalizeString, err)
	}
	lowercased := strings.ToLower(normalized)

	// Replace spaces with dashes.
	space := regexp.MustCompile(`\s+`)
	hyphened := space.ReplaceAllString(lowercased, "-")

	// Remove all non-alphanumeric characters except dashes.
	reg := regexp.MustCompile(`[^a-z0-9\-]+`)
	cleaned := reg.ReplaceAllString(hyphened, "")

	// Avoid double dashes caused by previous replacements.
	doubleDashes := regexp.MustCompile(`\-+`)
	slug := doubleDashes.ReplaceAllString(cleaned, "-")

	// Trim dashes at the beginning and end.
	slug = strings.Trim(slug, "-")

	return slug, nil
}
