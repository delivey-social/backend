package pedido

import "github.com/google/uuid"

type PedidoRepository interface {
	Create(
		items []PedidoItem,
		usuario Usuario,
		endereco Endereco,
		metodoPagamento PaymentMethods,
	) uuid.UUID

	FindByID(id uuid.UUID) (Pedido, error)

	UpdateStatus(id uuid.UUID, status PedidoStatus) error
}

type Pedido struct {
	ID            uuid.UUID
	Items         []PedidoItem
	Status        PedidoStatus
	Customer      Usuario
	Address       Endereco
	PaymentMethod PaymentMethods
}

type PedidoItem struct {
	ID            uuid.UUID
	Quantity      uint16
	PriceSnapshot uint32
}

type PedidoStatus string

const (
	PedidoStatusCreated          PedidoStatus = "CREATED"
	PedidoStatusReadyForDelivery PedidoStatus = "READY_FOR_DELIVERY"
	PedidoStatusInDelivery       PedidoStatus = "IN_DELIVERY"
	PedidoStatusDeliveryFinished PedidoStatus = "DELIVERED"
)
