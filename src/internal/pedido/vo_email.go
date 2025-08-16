package pedido

import (
	"errors"
	"net/mail"
	"regexp"
)

type Email struct {
	value string
}

var strictEmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-']+@([a-zA-Z0-9]+(-[a-zA-Z0-9]+)*\.)+[a-zA-Z]{2,63}$`)

func NewEmail(input string) (Email, error) {
	email, err := mail.ParseAddress(input)
	if err != nil {
		return Email{}, errors.New("email inválido")
	}

	if !strictEmailRegex.MatchString(email.Address) {
		return Email{}, errors.New("email inválido")
	}

	return Email{
		value: email.Address,
	}, nil
}
