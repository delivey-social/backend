package pedido

import "github.com/google/uuid"

type PedidoRepository interface {
	Create(
		items []PedidoItem,
		usuario Usuario,
		endereco Endereco,
		metodoPagamento PaymentMethod,
	) uuid.UUID

	FindByID(id uuid.UUID) (Pedido, error)

	UpdateStatus(id uuid.UUID, status PedidoStatus) error
}

type PedidoStatus string

const (
	PedidoStatusCreated          PedidoStatus = "CREATED"
	PedidoStatusReadyForDelivery PedidoStatus = "READY_FOR_DELIVERY"
	PedidoStatusInDelivery       PedidoStatus = "IN_DELIVERY"
	PedidoStatusDeliveryFinished PedidoStatus = "DELIVERED"
)
