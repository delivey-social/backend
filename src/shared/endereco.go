package shared

import "errors"

var (
	ErrInvalidCoordinates = errors.New("invalid latitude or longitude")
	ErrInvalidField       = errors.New("invalid field")
)

// TODO: Add tests
type Endereco struct {
	cep         CEP
	rua         string
	bairro      string
	numero      string
	complemento string
	cidade      string
	// TODO: Make UF it's own type (enum)
	uf string
	// TODO: Make coordinates VO
	latitude  float64
	longitude float64
}

func NewEndereco(
	cep CEP,
	rua string,
	bairro string,
	numero string,
	complemento string,
	cidade string,
	uf string,
	latitude float64,
	longitude float64) (Endereco, error) {

	if rua == "" || bairro == "" || numero == "" || cidade == "" || uf == "" {
		return Endereco{}, ErrInvalidField
	}

	// TODO: Limit latitude and longitude to Brazil limits
	if latitude < -90 || latitude > 90 {
		return Endereco{}, ErrInvalidCoordinates
	}
	if longitude < -180 || longitude > 180 {
		return Endereco{}, ErrInvalidCoordinates
	}

	return Endereco{
		cep:         cep,
		rua:         rua,
		bairro:      bairro,
		numero:      numero,
		complemento: complemento,
		cidade:      cidade,
		uf:          uf,
		latitude:    latitude,
		longitude:   longitude,
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
