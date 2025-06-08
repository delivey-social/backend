package shared

import "errors"

var (
	ErrInvalidCEP = errors.New("invalid CEP")
)

type Endereco struct {
	cep         CEP
	rua         string
	bairro      string
	numero      string
	complemento string
	cidade      string
	uf          string
	latitude    float64
	longitude   float64
}

func NewEndereco(cep CEP) (Endereco, error) {
	// TODO: Fetch address infos

	return Endereco{
		cep:         cep,
		rua:         "",
		bairro:      "",
		numero:      "",
		complemento: "",
		cidade:      "",
		uf:          "",
		latitude:    0,
		longitude:   0,
	}, nil
}

func (endereco *Endereco) CEP() CEP            { return endereco.cep }
func (endereco *Endereco) Rua() string         { return endereco.rua }
func (endereco *Endereco) Bairro() string      { return endereco.bairro }
func (endereco *Endereco) Numero() string      { return endereco.numero }
func (endereco *Endereco) Complemento() string { return endereco.complemento }
func (endereco *Endereco) Cidade() string      { return endereco.cidade }
func (endereco *Endereco) UF() string          { return endereco.uf }
func (endereco *Endereco) Latitude() float64   { return (endereco.latitude) }
func (endereco *Endereco) Longitude() float64  { return endereco.longitude }
