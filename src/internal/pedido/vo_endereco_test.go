package pedido_test

import (
	"testing"

	"comida.app/src/internal/pedido"
)

func TestEndereco(t *testing.T) {
	validCEP, _ := pedido.NewCEP("80000000")
	validRua := "Rua Teste"
	validBairro := "Bairro Teste"
	validNumero := "100"
	validComplemento := "Apto 1"
	// validCidade := "Cidade Teste"
	// validUF := "SP"

	t.Run("It should create an endereço in all valid values", func(t *testing.T) {
		endereco, err := pedido.NewEndereco(
			validCEP, validRua, validBairro, validNumero, validComplemento,
		)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if endereco.CEP != validCEP {
			t.Errorf("expected CEP %v, got %v", validCEP, endereco.CEP)
		}
		if endereco.Rua != validRua {
			t.Errorf("expected Rua %v, got %v", validRua, endereco.Rua)
		}
		if endereco.Bairro != validBairro {
			t.Errorf("expected Bairro %v, got %v", validBairro, endereco.Bairro)
		}
		if endereco.Numero != validNumero {
			t.Errorf("expected Numero %v, got %v", validNumero, endereco.Numero)
		}
		if endereco.Complemento != validComplemento {
			t.Errorf("expected Complemento %v, got %v", validComplemento, endereco.Complemento)
		}

	})

	t.Run("It should return an error in invalid values", func(t *testing.T) {
		_, err := pedido.NewEndereco(
			validCEP, "", validBairro, validNumero, validComplemento,
		)
		if err != pedido.ErrInvalidField {
			t.Errorf("expected ErrInvalidField, got %v", err)
		}
	})
}
