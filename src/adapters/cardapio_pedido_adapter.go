package adapters

import (
	"comida.app/src/internal/cardapio"
	"comida.app/src/internal/pedido"
	"github.com/google/uuid"
)

type CardapioPedidoAdapter struct {
	service *cardapio.CardapioService
}

func NewCardapioPedidoAdapter(service *cardapio.CardapioService) *CardapioPedidoAdapter {
	return &CardapioPedidoAdapter{
		service,
	}
}

func (adapter *CardapioPedidoAdapter) GetItemsByIDS(ids []uuid.UUID) ([]pedido.CardapioItem, error) {
	items := adapter.service.GetItemsByIDS(ids)

	result := make([]pedido.CardapioItem, len(items))
	for i, item := range items {
		result[i] = pedido.CardapioItem{
			Id:    item.ID,
			Price: item.Price,
		}
	}

	return result, nil
}
