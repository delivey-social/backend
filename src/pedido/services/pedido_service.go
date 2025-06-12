package services

import (
	"comida.app/src/pedido"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

type PedidoService struct {
	RestauranteService RestauranteService
	EnderecoService    EnderecoService
	PedidoRepository   PedidoRepository
}

type PedidoRepository interface {
	Save(pedido pedido.Pedido) error
}

type RestauranteService interface {
	GetItemsSnapshot(itemIds []uuid.UUID) ([]valueobject.ItemPedidoSnapshot, error)
	GetCoordinates() (shared.Coordenada, error)
}

type EnderecoService interface {
	CreateEnderecoFromCEP(cep shared.CEP) (shared.Endereco, error)
	CalculateDeliveryFee(restauranteCoordinates shared.Coordenada, pedidoCoordinates shared.Coordenada) (int, error)
}

func NewPedidoService(
	restauranteService RestauranteService,
	enderecoService EnderecoService,
	pedidoRepository PedidoRepository,
) PedidoService {
	return PedidoService{
		RestauranteService: restauranteService,
		EnderecoService:    enderecoService,
		PedidoRepository:   pedidoRepository,
	}
}
