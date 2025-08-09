package adapters

import (
	"comida.app/src/internal/pedido"
	"comida.app/src/internal/restaurante"
	"github.com/google/uuid"
)

type CardapioPedidoAdapter struct {
	service *restaurante.RestauranteService
}

func NewCardapioPedidoAdapter(service *restaurante.RestauranteService) *CardapioPedidoAdapter {
	return &CardapioPedidoAdapter{
		service,
	}
}

func (adapter *CardapioPedidoAdapter) GetItemsByIDS(ids []uuid.UUID) ([]pedido.CardapioItem, error) {
	items := adapter.service.GetMenuItemsByIDs(ids)

	result := make([]pedido.CardapioItem, len(items))
	for i, item := range items {
		result[i] = pedido.CardapioItem{
			Id:    item.ID,
			Price: item.Price,
		}
	}

	return result, nil
}
