package restaurante

import (
	"errors"
	"strings"

	"comida.app/src/restaurante/valueobject"
	"comida.app/src/shared"
)

type Restaurante struct {
	Nome        string
	CNPJ        valueobject.CNPJ
	Endereco    shared.Endereco
	Responsavel shared.Usuario
	ImagemUrl   shared.URL
}

var (
	ErrInvalidNameLength = errors.New("restaurant name must have at least 3 digits")
)

func NewRestaurante(nome string, cnpj valueobject.CNPJ, endereco shared.Endereco, responsavel shared.Usuario, imagemUrl shared.URL) (Restaurante, error) {
	nome = strings.TrimSpace(nome)
	if len(nome) < 3 {
		return Restaurante{}, ErrInvalidNameLength
	}

	return Restaurante{
		Nome:        nome,
		CNPJ:        cnpj,
		Endereco:    endereco,
		Responsavel: responsavel,
		ImagemUrl:   imagemUrl,
	}, nil
}
