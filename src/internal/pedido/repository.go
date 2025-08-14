package pedido

import "github.com/google/uuid"

type PedidoRepository interface {
	Create(items []PedidoItem) uuid.UUID

	FindByID(id uuid.UUID) (Pedido, error)

	ReadyForDelivery(id uuid.UUID) error
}

type Pedido struct {
	ID    uuid.UUID
	Items []PedidoItem
	Status PedidoStatus
}

type PedidoItem struct {
	ID            uuid.UUID
	Quantity      uint16
	PriceSnapshot uint32
}

type PedidoStatus string

var(
	PedidoStatusCreated PedidoStatus = "CREATED"
	PedidoStatusReadyForDelivery PedidoStatus ="READY_FOR_DELIVERY"
)