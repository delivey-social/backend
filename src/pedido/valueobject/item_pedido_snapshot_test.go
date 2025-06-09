package valueobject_test

import (
	"testing"

	"comida.app/src/pedido/valueobject"
	"github.com/google/uuid"
)

func TestItemPedidoSnapshot(t *testing.T) {
	t.Run("TestNewItemPedidoSnapshot_Success", func(t *testing.T) {
		id := uuid.New()
		nome := "Produto Teste"
		preco := 100
		quantidade := 2

		item, err := valueobject.NewItemPedidoSnapshot(id, nome, preco, quantidade)
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if item.ItemId != id {
			t.Errorf("expected ItemId %v, got %v", id, item.ItemId)
		}
		if item.Nome != nome {
			t.Errorf("expected Nome %v, got %v", nome, item.Nome)
		}
		if item.Preco != preco {
			t.Errorf("expected Preco %v, got %v", preco, item.Preco)
		}
		if item.Quantidade != quantidade {
			t.Errorf("expected Quantidade %v, got %v", quantidade, item.Quantidade)
		}
	})

	t.Run("TestNewItemPedidoSnapshot_InvalidPreco", func(t *testing.T) {
		id := uuid.New()
		nome := "Produto Teste"
		preco := 0
		quantidade := 2

		_, err := valueobject.NewItemPedidoSnapshot(id, nome, preco, quantidade)
		if err != valueobject.ErrInvalidItemsPrice {
			t.Errorf("expected ErrInvalidItemsPrice, got %v", err)
		}
	})

	t.Run("TestNewItemPedidoSnapshot_InvalidQuantidade", func(t *testing.T) {
		id := uuid.New()
		nome := "Produto Teste"
		preco := 100
		quantidade := 0

		_, err := valueobject.NewItemPedidoSnapshot(id, nome, preco, quantidade)
		if err != valueobject.ErrInvalidQuantity {
			t.Errorf("expected ErrInvalidQuantity, got %v", err)
		}
	})
}
