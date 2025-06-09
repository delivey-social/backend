package shared_test

import (
	"testing"

	"comida.app/src/shared"
)

func TestNewUrl(t *testing.T) {
	tests := []struct {
		input       string
		expectError bool
	}{
		{"https://example.com", false},
		{"http://localhost:8080/path", false},
		{"ftp://ftp.example.com", false},
		{"not a url", true},
		{"", true},
		{"http:/incomplete.com", true},
		{"http://", true},
	}

	for _, tt := range tests {
		urlObj, err := shared.NewUrl(tt.input)
		if tt.expectError {
			if err != shared.ErrInvalidURL {
				t.Errorf("NewUrl(%q) expected ErrInvalidURL, got %v", tt.input, err)
			}
		} else {
			if err != nil {
				t.Errorf("NewUrl(%q) unexpected error: %v", tt.input, err)
			}
			if urlObj.Value != tt.input {
				t.Errorf("NewUrl(%q) Value = %q, want %q", tt.input, urlObj.Value, tt.input)
			}
		}
	}
}
