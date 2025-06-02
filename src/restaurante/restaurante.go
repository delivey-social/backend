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
	ErrInvalidCNPJ = errors.New("invalid CNPJ")
)

func NewRestaurante(nome string, cnpj string, cep string, responsavel shared.Usuario, imagem_url string) (Restaurante, error) {
	endereco := shared.Endereco{
		CEP:         cep,
		Rua:         "",
		Bairro:      "",
		Numero:      "",
		Complemento: "",
		Cidade:      "",
		UF:          "",
		Latitude:    0,
		Longitude:   0,
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
