package shared

import (
	"errors"
	"regexp"
)

type CEP struct {
	value string
}

var (
	ErrInvalidFormat = errors.New("invalid CEP format")

	sanitizationRegex     = regexp.MustCompile(`[.-]`)
	digitsValidationRegex = regexp.MustCompile(`^\d{8}$`)
)

func NewCEP(input string) (CEP, error) {
	sanitized := sanitizationRegex.ReplaceAllString(input, "")

	if !digitsValidationRegex.MatchString(sanitized) {
		return CEP{}, ErrInvalidFormat
	}

	return CEP{
		value: sanitized,
	}, nil
}

func (cep CEP) String() string {
	return cep.value
}

func (cep CEP) Format() string {
	return cep.value[:2] + "." + cep.value[2:5] + "-" + cep.value[5:]
}
