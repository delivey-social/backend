package valueobject

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
