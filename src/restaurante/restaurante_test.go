package restaurante_test

import (
	"testing"

	"comida.app/src/restaurante"
	"comida.app/src/restaurante/valueobject"
	"comida.app/src/shared"
)

func makeValidRestauranteArgs() (string, valueobject.CNPJ, shared.Endereco, shared.Usuario, shared.URL) {
	cnpj, _ := valueobject.NewCNPJ("12345678000199")
	endereco := shared.Endereco{} // Assuming a zero value is valid for testing
	usuario := shared.Usuario{}   // Assuming a zero value is valid for testing
	url, _ := shared.NewUrl("https://imagem.com/foto.png")
	return "Restaurante Bom", cnpj, endereco, usuario, url
}

func TestRestaurante(t *testing.T) {

	t.Run("TestNewRestaurante_Success", func(t *testing.T) {
		nome, cnpj, endereco, usuario, url := makeValidRestauranteArgs()
		r, err := restaurante.NewRestaurante(nome, cnpj, endereco, usuario, url)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if r.Nome != nome {
			t.Errorf("expected nome %q, got %q", nome, r.Nome)
		}
	})

	t.Run("TestNewRestaurante_NameTooShort", func(t *testing.T) {
		_, cnpj, endereco, usuario, url := makeValidRestauranteArgs()
		_, err := restaurante.NewRestaurante("ab", cnpj, endereco, usuario, url)
		if err != restaurante.ErrInvalidNameLength {
			t.Errorf("expected ErrInvalidNameLength, got %v", err)
		}
	})

	t.Run("TestNewRestaurante_NameWithSpaces", func(t *testing.T) {
		_, cnpj, endereco, usuario, url := makeValidRestauranteArgs()
		inputName := "   Restaurante Legal   "
		expectedName := "Restaurante Legal"
		r, err := restaurante.NewRestaurante(inputName, cnpj, endereco, usuario, url)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if r.Nome != expectedName {
			t.Errorf("expected nome %q, got %q", expectedName, r.Nome)
		}
	})
}
