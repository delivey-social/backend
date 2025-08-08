package pedido

import (
	"github.com/google/uuid"
)

type PedidoService struct {
	repository      PedidoRepository
	cardapioService CardapioService
}

func NewPedidoService(repository PedidoRepository, cardapioService CardapioService) *PedidoService {
	return &PedidoService{
		repository,
		cardapioService,
	}
}

func (s *PedidoService) Create(items []CreatePedidoRequestItem) error {
	var itemsIDs []uuid.UUID
	for _, item := range items {
		itemsIDs = append(itemsIDs, item.ItemID)
	}

	// Check if items exist
	menuItems, err := s.cardapioService.GetItemsByIDS(itemsIDs)
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
