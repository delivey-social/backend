package pedido

import (
	"errors"
	"regexp"
)

type Telefone struct {
	countryCode string
	ufCode      string
	phoneNumber string
}

var (
	validateRegex  = regexp.MustCompile(`^(\+?\s?55)?\s?(\(?\d{2}\)?\s?)\d{4,5}(-?)\d{4}$`)
	normalizeRegex = regexp.MustCompile(`^\+?\s?55|\D`)
)

func NewTelefone(input string) (Telefone, error) {
	if !validateRegex.MatchString(input) {
		return Telefone{}, errors.New("telefone inválido")
	}

	sanitized := normalizeRegex.ReplaceAllString(input, "")

	ufCode := sanitized[:2]
	phone := sanitized[2:]

	return Telefone{
		countryCode: "55",
		ufCode:      ufCode,
		phoneNumber: phone,
	}, nil
}

func (telefone Telefone) String() string {
	return telefone.countryCode + telefone.ufCode + telefone.phoneNumber
}

func (telefone Telefone) Formatted() string {
	raw := telefone.String()

	re := regexp.MustCompile(`^(\d{2})(\d{2})(\d{4,5})(\d{4})`)
	formatted := re.ReplaceAllString(raw, "+$1 ($2) $3-$4")

	return formatted
}
