package pedido

import (
	"errors"
)

var (
	ErrInvalidField = errors.New("invalid field")
)

type Endereco struct {
	CEP         CEP
	Rua         string
	Bairro      Bairro
	Numero      string
	Complemento string
	// Cidade      string
	// TODO: Make UF it's own type (enum)
	// UF string
}

func NewEndereco(
	cep CEP,
	rua string,
	bairro Bairro,
	numero string,
	complemento string) (Endereco, error) {

	if rua == "" || numero == "" {
		return Endereco{}, ErrInvalidField
	}

	return Endereco{
		CEP:         cep,
		Rua:         rua,
		Bairro:      bairro,
		Numero:      numero,
		Complemento: complemento,
		// Cidade:      cidade,
		// UF:          uf,
	}, nil
}
