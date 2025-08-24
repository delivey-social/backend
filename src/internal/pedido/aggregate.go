package pedido

import "github.com/google/uuid"

type Pedido struct {
	ID            uuid.UUID
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

func (p *Pedido) CalculateTotal() uint32 {
	var itemsTotal uint32

	for _, item := range p.Items {
		itemsTotal += item.PriceSnapshot * uint32(item.Quantity)
	}

	return itemsTotal
}
