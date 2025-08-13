package pedido

import (
	"errors"

	"comida.app/src/infra"
	"github.com/google/uuid"
)

type PedidoService struct {
	repository      PedidoRepository
	cardapioService RestauranteService
	publisher EventPublisher
}

func NewPedidoService(repository PedidoRepository, cardapioService RestauranteService, publisher EventPublisher) *PedidoService {
	return &PedidoService{
		repository,
		cardapioService,
		publisher,
	}
}

func (s *PedidoService) Create(restaurantID uuid.UUID, items []CreatePedidoRequestItem) (uuid.UUID, error) {
	if len(items) == 0 {
		return uuid.UUID{}, errors.New("é necessário que o pedido tenha ao menos um item")
	}

	for _, item := range items {
		if item.Quantity <= 0 {
			return uuid.UUID{}, errors.New("algum item possuí quantidade inválida")
		}
	}

	var itemsIDs []uuid.UUID
	for _, item := range items {
		itemsIDs = append(itemsIDs, item.ItemID)
	}

	menuItems, err := s.cardapioService.GetItemsByIDS(restaurantID, itemsIDs)
	if err != nil {
		return uuid.UUID{}, err
	}

	// Creates the pedido
	id := s.repository.Create(joinItems(items, menuItems))

	s.publisher.Publish(infra.OrderCreated)

	return id, nil
}

func joinItems(quantities []CreatePedidoRequestItem, prices []CardapioItem) []PedidoItem {
	priceMap := make(map[string]uint32)

	for _, price := range prices {
		priceMap[price.Id.String()] = price.Price
	}

	var result []PedidoItem
	for _, item := range quantities {
		price, ok := priceMap[item.ItemID.String()]
		if !ok {
			panic("item not found")
		}

		result = append(result, PedidoItem{
			ID:            item.ItemID,
			Quantity:      item.Quantity,
			PriceSnapshot: price,
		})
	}

	return result
}
