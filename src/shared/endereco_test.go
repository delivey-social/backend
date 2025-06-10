package shared_test

import (
	"testing"

	"comida.app/src/shared"
)

func TestEndereco(t *testing.T) {
	validCEP, _ := shared.NewCEP("80000000")
	validRua := "Rua Teste"
	validBairro := "Bairro Teste"
	validNumero := "100"
	validComplemento := "Apto 1"
	validCidade := "Cidade Teste"
	validUF := "SP"
	validCoordinate, _ := shared.NewCoordenada(50.44, 40.44)

	t.Run("It should create an endereço in all valid values", func(t *testing.T) {
		endereco, err := shared.NewEndereco(
			validCEP, validRua, validBairro, validNumero, validComplemento,
			validCidade, validUF, validCoordinate,
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
		if endereco.Cidade != validCidade {
			t.Errorf("expected Cidade %v, got %v", validCidade, endereco.Cidade)
		}
		if endereco.UF != validUF {
			t.Errorf("expected UF %v, got %v", validUF, endereco.UF)
		}
		if endereco.Coordenada != validCoordinate {
			t.Errorf("expected Coordenada %v, got %v", validCoordinate, endereco.Coordenada)
		}
	})

	t.Run("It should return an error in invalid values", func(t *testing.T) {
		_, err := shared.NewEndereco(
			validCEP, "", validBairro, validNumero, validComplemento,
			validCidade, validUF, validCoordinate,
		)
		if err != shared.ErrInvalidField {
			t.Errorf("expected ErrInvalidField, got %v", err)
		}
	})
}
