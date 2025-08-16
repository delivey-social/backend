package pedido_test

import (
	"testing"

	"comida.app/src/internal/pedido"
)

func TestNewUsuario(t *testing.T) {
	email, err := pedido.NewEmail("user@example.com")
	if err != nil {
		t.Fatal(err)
	}
	telefone, err := pedido.NewTelefone("11999999999")
	if err != nil {
		t.Fatal(err)
	}

	nome := "Test"

	usuario := pedido.NewUsuario(email, telefone, nome)

	if usuario.Email() != email {
		t.Errorf("expected email %v, got %v", email, usuario.Email())
	}

	if usuario.Telefone() != telefone {
		t.Errorf("expected telefone %v, got %v", telefone, usuario.Telefone())
	}
}
