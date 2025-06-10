package shared_test

import (
	"testing"

	"comida.app/src/shared"
)

func TestNewCoordenada_ValidCoordinates(t *testing.T) {
	tests := []struct {
		name      string
		latitude  float32
		longitude float32
	}{
		{"Zero coordinates", 0, 0},
		{"Positive coordinates", 10.5, 20.5},
		{"Negative coordinates", -10.5, -20.5},
		{"Max latitude", 90, 0},
		{"Min latitude", -90, 0},
		{"Max longitude", 0, 180},
		{"Min longitude", 0, -180},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coord, err := shared.NewCoordenada(tt.latitude, tt.longitude)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}
			if coord.Latitude != tt.latitude || coord.Longitude != tt.longitude {
				t.Errorf("expected (%v, %v), got (%v, %v)", tt.latitude, tt.longitude, coord.Latitude, coord.Longitude)
			}
		})
	}
}

func TestNewCoordenada_InvalidLatitude(t *testing.T) {
	_, err := shared.NewCoordenada(91, 0)
	if err != shared.ErrInvalidCoordinates {
		t.Errorf("expected ErrInvalidCoordinates for latitude > 90, got %v", err)
	}
	_, err = shared.NewCoordenada(-91, 0)
	if err != shared.ErrInvalidCoordinates {
		t.Errorf("expected ErrInvalidCoordinates for latitude < -90, got %v", err)
	}
}

func TestNewCoordenada_InvalidLongitude(t *testing.T) {
	_, err := shared.NewCoordenada(0, 181)
	if err != shared.ErrInvalidCoordinates {
		t.Errorf("expected ErrInvalidCoordinates for longitude > 180, got %v", err)
	}
	_, err = shared.NewCoordenada(0, -181)
	if err != shared.ErrInvalidCoordinates {
		t.Errorf("expected ErrInvalidCoordinates for longitude < -180, got %v", err)
	}
}
