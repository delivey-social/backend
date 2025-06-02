package shared

type Endereco struct {
	// TODO: CEP should be its own value object
	CEP         string
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
	// TODO: Fetch address infos in CEP

	return Endereco{
		CEP:         cep,
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
