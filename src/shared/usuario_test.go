package shared_test

import (
	"testing"

	"comida.app/src/shared"
)

func TestNewUsuario(t *testing.T) {
	email, err := shared.NewEmail("user@example.com")
	if err != nil {
		t.Fatal(err)
	}
	telefone, err := shared.NewTelefone("11999999999")
	if err != nil {
		t.Fatal(err)
	}

	usuario := shared.NewUsuario(email, telefone)

	if usuario.Email() != email {
		t.Errorf("expected email %v, got %v", email, usuario.Email())
	}

	if usuario.Telefone() != telefone {
		t.Errorf("expected telefone %v, got %v", telefone, usuario.Telefone())
	}
}
