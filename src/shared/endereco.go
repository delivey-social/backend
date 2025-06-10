package shared

import "errors"

var (
	ErrInvalidField = errors.New("invalid field")
)

type Endereco struct {
	CEP         CEP
	Rua         string
	Bairro      string
	Numero      string
	Complemento string
	Cidade      string
	// TODO: Make UF it's own type (enum)
	UF         string
	Coordenada Coordenada
}

func NewEndereco(
	cep CEP,
	rua string,
	bairro string,
	numero string,
	complemento string,
	cidade string,
	uf string,
	coordenada Coordenada) (Endereco, error) {

	if rua == "" || bairro == "" || numero == "" || cidade == "" || uf == "" {
		return Endereco{}, ErrInvalidField
	}

	return Endereco{
		CEP:         cep,
		Rua:         rua,
		Bairro:      bairro,
		Numero:      numero,
		Complemento: complemento,
		Cidade:      cidade,
		UF:          uf,
		Coordenada:  coordenada,
	}, nil
}
