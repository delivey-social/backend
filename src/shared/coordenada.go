package shared

import "errors"

type Coordenada struct {
	Latitude  float32
	Longitude float32
}

var (
	ErrInvalidCoordinates = errors.New("invalid latitude or longitude")
)

func NewCoordenada(latitude float32, longitude float32) (Coordenada, error) {
	// TODO: Limit latitude and longitude to Brazil limits
	if latitude < -90 || latitude > 90 {
		return Coordenada{}, ErrInvalidCoordinates
	}
	if longitude < -180 || longitude > 180 {
		return Coordenada{}, ErrInvalidCoordinates
	}

	return Coordenada{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
