package pedido

import (
	"errors"

	"github.com/google/uuid"
)

type PedidoService struct {
	repository      PedidoRepository
	cardapioService RestauranteService
}

func NewPedidoService(repository PedidoRepository, cardapioService RestauranteService) *PedidoService {
	return &PedidoService{
		repository,
		cardapioService,
	}
}

func (s *PedidoService) Create(restaurantID uuid.UUID, items []CreatePedidoRequestItem) error {
	if len(items) == 0 {
		return errors.New("é necessário que o pedido tenha ao menos um item")
	}

	for _, item := range items {
		if item.Quantity <= 0 {
			return errors.New("algum item possuí quantidade inválida")
		}
	}

	var itemsIDs []uuid.UUID
	for _, item := range items {
		itemsIDs = append(itemsIDs, item.ItemID)
	}

	// Check if items exist
	menuItems, err := s.cardapioService.GetItemsByIDS(restaurantID, itemsIDs)
	if err != nil {
		return err
	}

	// Creates the pedido
	s.repository.Create(joinItems(items, menuItems))

	return nil
}

func joinItems(quantities []CreatePedidoRequestItem, prices []CardapioItem) []PedidoRepositoryItem {
	priceMap := make(map[string]uint32)

	for _, price := range prices {
		priceMap[price.Id.String()] = price.Price
	}

	var result []PedidoRepositoryItem
	for _, item := range quantities {
		price, ok := priceMap[item.ItemID.String()]
		if !ok {
			panic("item not found")
		}

		result = append(result, PedidoRepositoryItem{
			ID:            item.ItemID,
			Quantity:      item.Quantity,
			PriceSnapshot: price,
		})
	}

	return result
}
