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

func (adapter *CardapioPedidoAdapter) GetItemsByIDS(restaurantID uuid.UUID, ids []uuid.UUID) ([]pedido.CardapioItem, error) {
	items, err := adapter.service.GetMenuItemsByIDs(restaurantID, ids)
	if err != nil {
		return []pedido.CardapioItem{}, err
	}

	result := make([]pedido.CardapioItem, len(*items))
	for i, item := range *items {
		result[i] = pedido.CardapioItem{
			Id:    item.ID,
			Price: item.Price,
		}
	}

	return result, nil
}
