package stringx_test

import (
	"testing"

	"github.com/dmitrymomot/go-app-template/pkg/stringx"
)

func TestToSlug(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Hello World",
			expected: "hello-world",
		},
		{
			input:    "This is a Test",
			expected: "this-is-a-test",
		},
		{
			input:    "12345",
			expected: "12345",
		},
		{
			input:    "Special Characters: !@#$%^&*()_+",
			expected: "special-characters",
		},
		{
			input:    "This is a test with numbers 12345",
			expected: "this-is-a-test-with-numbers-12345",
		},
		{
			input:    "This is a test with special characters: !@#$%^&*()_+",
			expected: "this-is-a-test-with-special-characters",
		},
		{
			input:    "thisIsATest",
			expected: "thisisatest",
		},
	}

	for _, test := range tests {
		result, err := stringx.ToSlug(test.input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Got: %s", test.input, test.expected, result)
		}
	}
}
