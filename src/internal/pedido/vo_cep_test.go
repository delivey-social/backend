package pedido_test

import (
	"testing"

	"comida.app/src/internal/pedido"
)

var validCeps = []string{"80230000", "80.230-000", "80.230000", "80230-000"}
var invalidCeps = []string{"8023000", "80.230-000a", "80.2300000"}

type CPFValues struct {
	Input     string
	Value     string
	Formatted string
}

var mockCpf = CPFValues{
	Input:     "80.230000",
	Value:     "80230000",
	Formatted: "80.230-000",
}

func TestCep(t *testing.T) {
	t.Run("valid CEPs", func(t *testing.T) {
		t.Parallel()
		for _, cep := range validCeps {
			_, err := pedido.NewCEP(cep)
			if err != nil {
				t.Errorf("Unexpected error %v", err)
			}
		}
	})

	t.Run("Invalid CEPs", func(t *testing.T) {
		t.Parallel()
		for _, cep := range invalidCeps {
			_, err := pedido.NewCEP(cep)
			if err == nil {
				t.Errorf("Expected error, got nothing %s", cep)
			}
		}
	})

	t.Run("cpf.String", func(t *testing.T) {
		t.Parallel()
		cepObj, err := pedido.NewCEP(mockCpf.Input)
		if err != nil {
			t.Errorf("Unexpected error %v", err)
		}
		if cepObj.String() != mockCpf.Value {
			t.Errorf("Expected sanitized CEP, got %s", cepObj.String())
		}
	})

	t.Run("cpf.Format", func(t *testing.T) {
		t.Parallel()
		cepObj, err := pedido.NewCEP(mockCpf.Input)
		if err != nil {
			t.Fatal("Unexpected Error")
		}

		if cepObj.Format() != mockCpf.Formatted {
			t.Errorf("Expected formatted CEP, got %s", cepObj.Format())
		}
	})
}
