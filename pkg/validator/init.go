package validator

import (
	"fmt"

	"braces.dev/errtrace"
	"github.com/dmitrymomot/go-app-template/pkg/emailx"
	"github.com/gookit/validate"
	"github.com/pkg/errors"
)

func init() {
	// change global opts
	validate.Config(func(opt *validate.GlobalOption) {
		opt.FieldTag = "form"
		opt.StopOnError = false
		opt.SkipOnEmpty = true
		opt.UpdateSource = true
		opt.CheckZero = false
		opt.ErrKeyFmt = 1
	})

	// Add custom global validation rules
	validate.AddValidators(validate.M{
		"realEmail": func(val interface{}) bool {
			email, ok := val.(string)
			if !ok {
				return false
			}
			if _, err := emailx.SanitizeEmail(email, true); err != nil {
				return false
			}
			return true
		},
		"password": func(val interface{}) bool {
			password, ok := val.(string)
			if !ok {
				return false
			}
			if err := ValidatePassword(password); err != nil {
				return false
			}
			return true
		},
	})

	// Add global filters
	validate.AddFilters(validate.M{
		"sanitizeEmail": func(val interface{}) (string, error) {
			if email, ok := val.(string); ok {
				return errtrace.Wrap2(emailx.SanitizeEmail(email, false))
			}
			return fmt.Sprintf("%v", val), errtrace.Wrap(errors.New("invalid email address"))
		},
	})

	// Add global messages
	validate.AddGlobalMessages(map[string]string{
		"realEmail":     "Email address seems not real",
		"sanitizeEmail": "Invalid email address",
		"password":      "Password must contain at least one digit, uppercase, and lowercase letters",
	})
}
