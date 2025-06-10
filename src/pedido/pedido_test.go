package pedido_test

import (
	"testing"

	"comida.app/src/pedido"
	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

func makeValidUsuario() shared.Usuario {
	email, _ := shared.NewEmail("test@test.com")
	telefone, _ := shared.NewTelefone("41 99999-9999")

	return shared.NewUsuario(email, telefone)
}

func makeValidEndereco() shared.Endereco {
	cep, _ := shared.NewCEP("80000000")
	coordenada, _ := shared.NewCoordenada(50.65, 40.44)

	endereco, _ := shared.NewEndereco(cep, "Rua de Teste", "Teste", "994", "AP 12", "Curitiba", "PR", coordenada)

	return endereco
}

func makeValidPreco() valueobject.Preco {
	preco, _ := valueobject.NewPreco(100, 10, 5)

	return preco
}

func makeValidItemPedidoSnapshot() valueobject.ItemPedidoSnapshot {
	itemId := uuid.New()
	item, _ := valueobject.NewItemPedidoSnapshot(itemId, "Produto", 50, 2)
	return item
}

func TestNewPedido_Success(t *testing.T) {
	cliente := makeValidUsuario()
	endereco := makeValidEndereco()
	preco := makeValidPreco()
	item := makeValidItemPedidoSnapshot()
	itens := []valueobject.ItemPedidoSnapshot{item}
	observacao := "Sem cebola"
	metodo := enums.Pix

	p, err := pedido.NewPedido(cliente, itens, endereco, preco, observacao, metodo)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if p.Cliente.Email() != cliente.Email() {
		t.Errorf("expected cliente email %v, got %v", cliente.Email(), p.Cliente.Email())
	}
	if len(p.Itens) != 1 {
		t.Errorf("expected 1 item, got %d", len(p.Itens))
	}
	if p.Preco.Total() != preco.Total() {
		t.Errorf("expected preco total %d, got %d", preco.Total(), p.Preco.Total())
	}
	if p.Observacao != observacao {
		t.Errorf("expected observacao %q, got %q", observacao, p.Observacao)
	}
	if p.Status != enums.StatusAguardandoPagamento {
		t.Errorf("expected status %v, got %v", enums.StatusAguardandoPagamento, p.Status)
	}
	if p.MetodoPagamento != metodo {
		t.Errorf("expected metodo_pagamento %v, got %v", metodo, p.MetodoPagamento)
	}
}
