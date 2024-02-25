package validator

import (
	"net/url"

	"github.com/gookit/validate"
)

// ValidateStruct validate struct data
// - s: struct pointer
// - return: validation errors as url.Values or nil if no errors
func ValidateStruct(s interface{}) url.Values {
	v := validate.Struct(s)
	if v.Validate() {
		return nil
	}

	if !v.Errors.Empty() {
		// cast errors to url.Values
		result := url.Values{}
		for k, errs := range v.Errors.All() {
			for _, err := range errs {
				result.Add(k, err)
			}
		}

		return result
	}

	return nil
}
