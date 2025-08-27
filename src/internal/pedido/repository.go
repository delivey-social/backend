package pedido

import "github.com/google/uuid"

type PedidoRepository interface {
	Create(pedido Pedido)
	Update(pedido Pedido) error

	FindByID(id uuid.UUID) (Pedido, error)
}

type PedidoStatus string

const (
	PedidoStatusCreated          PedidoStatus = "CREATED"
	PedidoStatusReadyForDelivery PedidoStatus = "READY_FOR_DELIVERY"
	PedidoStatusInDelivery       PedidoStatus = "IN_DELIVERY"
	PedidoStatusDeliveryFinished PedidoStatus = "DELIVERED"
)
