package valueobject

import (
	"errors"
	"fmt"
	"regexp"
)

type CNPJ struct {
	value string
}

var (
	ErrInvalidFormat            = errors.New("invalid format")
	ErrInvalidVerificationDigit = errors.New("invalid verification digit")
	ErrAllDigitsEqual           = errors.New("all digits are equal")

	validateRegex = regexp.MustCompile(`\d{2}\.?\d{3}\.?\d{3}\/?\d{4}\-?\d{2}`)
	replaceRegex  = regexp.MustCompile(`[/.-]`)
)

func NewCNPJ(input string) (CNPJ, error) {
	if !validateRegex.MatchString(input) {
		return CNPJ{}, ErrInvalidFormat
	}

	sanitized := replaceRegex.ReplaceAllString(input, "")

	if areAllDigitsEqual(sanitized) {
		return CNPJ{}, ErrAllDigitsEqual
	}
	if !validateFirstDigit(sanitized) {
		return CNPJ{}, ErrInvalidVerificationDigit
	}
	if !validateSecondDigit(sanitized) {
		return CNPJ{}, ErrInvalidVerificationDigit
	}

	return CNPJ{
		value: sanitized,
	}, nil
}

func (cnpj CNPJ) String() string {
	return cnpj.value
}

func (cnpj CNPJ) Formatted() string {
	raw := cnpj.String()

	re := regexp.MustCompile(`^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})`)
	formatted := re.ReplaceAllString(raw, "$1.$2.$3/$4-$5")

	return formatted
}

func validateFirstDigit(cnpj string) bool {
	var validationDigit = int(cnpj[12] - '0')
	var validate = reverse(cnpj[:12])

	var sum int = 0
	for i, digit := range validate {
		digit := int(digit - '0')

		multiplier := 9 - i
		if multiplier < 2 {
			multiplier = 17 - i
		}
		sum += digit * multiplier
	}

	var validatedDigit = sum % 11
	if validatedDigit == 10 {
		validatedDigit = 0
	}
	if validationDigit != validatedDigit {
		fmt.Printf("Expected %d got %d\n", validationDigit, validatedDigit)
	}

	return validationDigit == validatedDigit
}
func validateSecondDigit(cnpj string) bool {
	var validationDigit = int(cnpj[13] - '0')
	var validate = reverse(cnpj[:13])

	var sum int = 0
	for i, digit := range validate {
		digit := int(digit - '0')

		multiplier := 9 - i
		if multiplier < 2 {
			multiplier = 17 - i
		}
		sum += digit * multiplier
	}

	var validatedDigit = sum % 11
	if validatedDigit == 10 {
		validatedDigit = 0
	}

	if validationDigit != validatedDigit {
		fmt.Printf("Expected %d got %d\n", validationDigit, validatedDigit)
	}

	return validationDigit == validatedDigit
}

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func areAllDigitsEqual(s string) bool {
	for i := range s {
		if i == 0 {
			continue
		}
		if s[i] != s[i-1] {
			return false
		}
	}

	return true
}
