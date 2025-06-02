package restaurante

import (
	"errors"

	"comida.app/src/restaurante/valueobject"
	"comida.app/src/shared"
)

type Restaurante struct {
	nome        string
	cnpj        valueobject.CNPJ
	endereco    shared.Endereco
	responsavel shared.Usuario
	imagem_url  string
}

var (
	ErrInvalidCNPJ    = errors.New("invalid CNPJ")
	ErrInvalidAddress = errors.New("invalid address")
)

func NewRestaurante(nome string, cnpj string, cep string, responsavel shared.Usuario, imagem_url string) (Restaurante, error) {
	endereco, err := shared.NewEndereco(cep)
	if err != nil {
		return Restaurante{}, ErrInvalidAddress
	}

	cnpj_vo, err := valueobject.NewCNPJ(cnpj)
	if err != nil {
		return Restaurante{}, ErrInvalidCNPJ
	}

	return Restaurante{
		nome:        nome,
		cnpj:        cnpj_vo,
		endereco:    endereco,
		responsavel: responsavel,
		imagem_url:  imagem_url,
	}, nil
}
