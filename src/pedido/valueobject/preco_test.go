package valueobject_test

import (
	"testing"

	"comida.app/src/pedido/valueobject"
)

func TestPrecoTotal(t *testing.T) {
	tests := []struct {
		name        string
		precoItens  int
		taxaEntrega int
		taxaApp     int
		expected    int
	}{
		{
			name:        "positive values",
			precoItens:  100,
			taxaEntrega: 20,
			taxaApp:     10,
			expected:    130,
		},
		{
			name:        "negative values",
			precoItens:  -50,
			taxaEntrega: 10,
			taxaApp:     5,
			expected:    -35,
		},
		{
			name:        "mixed values",
			precoItens:  200,
			taxaEntrega: -50,
			taxaApp:     0,
			expected:    150,
		},
	}

	errorTests := []struct {
		name        string
		precoItens  int
		taxaEntrega int
		taxaApp     int
		expected    error
	}{
		{
			name:        "precoItens equals zero",
			precoItens:  0,
			taxaEntrega: 500,
			taxaApp:     800,
			expected:    valueobject.ErrInvalidItemsPrice,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preco, err := valueobject.NewPreco(tt.precoItens, tt.taxaEntrega, tt.taxaApp)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			total := preco.Total()
			if total != tt.expected {
				t.Errorf("Preco.Total() = %d, want %d", total, tt.expected)
			}
		})
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := valueobject.NewPreco(tt.precoItens, tt.taxaEntrega, tt.taxaApp)
			if err == nil {
				t.Errorf("Expected error, got nothing")
			}
		})
	}
}
