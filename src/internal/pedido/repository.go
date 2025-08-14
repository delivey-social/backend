package pedido

import "github.com/google/uuid"

type PedidoRepository interface {
	Create(items []PedidoItem) uuid.UUID

	FindByID(id uuid.UUID) (Pedido, error)
}

type Pedido struct {
	ID    uuid.UUID
	Items []PedidoItem
}

type PedidoItem struct {
	ID            uuid.UUID
	Quantity      uint16
	PriceSnapshot uint32
}
