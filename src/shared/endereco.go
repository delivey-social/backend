package shared

import "errors"

var (
	ErrInvalidCEP = errors.New("invalid CEP")
)

type Endereco struct {
	CEP         CEP
	Rua         string
	Bairro      string
	Numero      string
	Complemento string
	Cidade      string
	UF          string
	Latitude    float64
	Longitude   float64
}

func NewEndereco(cep string) (Endereco, error) {
	CEP, err := NewCEP(cep)
	if err != nil {
		return Endereco{}, ErrInvalidCEP
	}

	// TODO: Fetch address infos with CEP

	return Endereco{
		CEP:         CEP,
		Rua:         "",
		Bairro:      "",
		Numero:      "",
		Complemento: "",
		Cidade:      "",
		UF:          "",
		Latitude:    0,
		Longitude:   0,
	}, nil
}
