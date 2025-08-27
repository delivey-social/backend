package pedido

import (
	"fmt"

	"github.com/google/uuid"
)

type Pedido struct {
	id            uuid.UUID
	Items         []PedidoItem
	Status        PedidoStatus
	Customer      Usuario
	Address       Endereco
	PaymentMethod PaymentMethod
}

type PedidoItem struct {
	ID            uuid.UUID
	Quantity      uint16
	PriceSnapshot uint32
}

func NewPedido(items []PedidoItem, customer Usuario, Address Endereco, paymentMethod PaymentMethod) Pedido {
	return Pedido{
		id:            uuid.New(),
		Items:         items,
		Status:        PedidoStatusCreated,
		Customer:      customer,
		Address:       Address,
		PaymentMethod: paymentMethod,
	}
}

func (p *Pedido) CalculateTotal() uint32 {
	var itemsTotal uint32

	for _, item := range p.Items {
		itemsTotal += item.PriceSnapshot * uint32(item.Quantity)
	}

	return itemsTotal
}

func (p *Pedido) GetId() uuid.UUID {
	return p.id
}

func (p *Pedido) UpdateStatus(newStatus PedidoStatus) error {
	return fmt.Errorf("not implemented")
}
