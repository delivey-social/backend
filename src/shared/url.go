package shared

import (
	"errors"
	"net/url"
	"regexp"
)

type URL struct {
	Value string
}

var (
	ErrInvalidURL = errors.New("invalid URL")

	validateURL = regexp.MustCompile(`^[a-zA-Z]+://[a-zA-Z]+`)
)

func NewUrl(input string) (URL, error) {
	_, err := url.ParseRequestURI(input)
	if err != nil {
		return URL{}, ErrInvalidURL
	}

	if !validateURL.MatchString(input) {
		return URL{}, ErrInvalidURL
	}

	return URL{
		Value: input,
	}, nil
}
