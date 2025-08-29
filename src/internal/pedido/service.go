package pedido

import (
	"errors"

	"comida.app/src/infra"
	"comida.app/src/internal/pedido/bairro"
	"github.com/google/uuid"
)

type PedidoService struct {
	repository      PedidoRepository
	cardapioService RestauranteService
	bairroService   bairro.BairroService
	publisher       EventPublisher
}

func NewPedidoService(repository PedidoRepository, cardapioService RestauranteService, bairroService bairro.BairroService, publisher EventPublisher) *PedidoService {
	return &PedidoService{
		repository,
		cardapioService,
		bairroService,
		publisher,
	}
}

func (s *PedidoService) Create(
	restaurantID uuid.UUID,
	items []CreatePedidoDTOItem,
	usuario Usuario,
	endereco Endereco,
	metodoPagamento PaymentMethod,
) (uuid.UUID, error) {
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
	pedido := NewPedido(joinItems(items, menuItems), usuario, endereco, metodoPagamento)
	id := pedido.GetId()
	s.repository.Save(pedido)

	s.publisher.Publish(infra.Event{
		Type: infra.OrderCreated,
		Payload: infra.OrderCreatedPayload{
			OrderID: id,
		},
	})

	return id, nil
}

func (s *PedidoService) ReadyForDelivery(id uuid.UUID) error {
	pedido, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	err = pedido.UpdateStatus(PedidoStatusReadyForDelivery)
	if err != nil {
		return err
	}

	s.repository.Update(pedido)

	s.publisher.Publish(infra.Event{
		Type: infra.OrderReadyForDelivery,
		Payload: infra.OrderUpdatedPayload{
			OrderID: id,
		},
	})

	return nil
}

func (s *PedidoService) InitiateDelivery(id uuid.UUID) error {
	pedido, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	err = pedido.UpdateStatus(PedidoStatusInDelivery)
	if err != nil {
		return err
	}

	s.repository.Update(pedido)

	s.publisher.Publish(infra.Event{
		Type: infra.OrderInDelivery,
		Payload: infra.OrderUpdatedPayload{
			OrderID: id,
		},
	})

	return nil
}

func (s *PedidoService) FinishDelivery(id uuid.UUID) error {
	pedido, err := s.repository.FindByID(id)
	if err != nil {
		return err
	}

	err = pedido.UpdateStatus(PedidoStatusDeliveryFinished)
	if err != nil {
		return err
	}

	s.repository.Update(pedido)

	s.publisher.Publish(infra.Event{
		Type: infra.OrderDelivered,
		Payload: infra.OrderUpdatedPayload{
			OrderID: id,
		},
	})

	return nil
}

func (s *PedidoService) FindByID(id uuid.UUID) (Pedido, error) {
	return s.repository.FindByID(id)
}

func joinItems(quantities []CreatePedidoDTOItem, prices []CardapioItem) []PedidoItem {
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
