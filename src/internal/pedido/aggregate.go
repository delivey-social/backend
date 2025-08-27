package pedido

import (
	"fmt"
	"math"

	"github.com/google/uuid"
	"golang.org/x/exp/slices"
)

type Pedido struct {
	id                    uuid.UUID
	Items                 []PedidoItem
	Status                PedidoStatus
	Customer              Usuario
	Address               Endereco
	PaymentMethod         PaymentMethod
	DeliveryFeeCalculator DeliveryFeeCalculator
}

type PedidoItem struct {
	ID            uuid.UUID
	Quantity      uint16
	PriceSnapshot uint32
}

type PedidoTotal struct {
	Itens       uint32 `json:"itens"`
	TaxaApp     uint32 `json:"taxa_aplicativo"`
	TaxaEntrega uint32 `json:"taxa_entrega"`
}

func NewPedido(items []PedidoItem, customer Usuario, Address Endereco, paymentMethod PaymentMethod) Pedido {
	return Pedido{
		id:                    uuid.New(),
		Items:                 items,
		Status:                PedidoStatusCreated,
		Customer:              customer,
		Address:               Address,
		PaymentMethod:         paymentMethod,
		DeliveryFeeCalculator: NewFixedRateCalculator(500),
	}
}

const APP_FEE_PERCENTAGE = float64(10) / 100

func (p *Pedido) CalculateTotal() PedidoTotal {
	var itemsTotal uint32

	for _, item := range p.Items {
		itemsTotal += item.PriceSnapshot * uint32(item.Quantity)
	}

	return PedidoTotal{
		Itens:       itemsTotal,
		TaxaApp:     uint32(math.Round(float64(itemsTotal) * APP_FEE_PERCENTAGE)),
		TaxaEntrega: p.DeliveryFeeCalculator.Calculate(DeliveryFeeContext{}),
	}
}

func (p *Pedido) GetId() uuid.UUID {
	return p.id
}

var validNextStatus map[PedidoStatus][]PedidoStatus = map[PedidoStatus][]PedidoStatus{
	PedidoStatusCreated:          {PedidoStatusReadyForDelivery},
	PedidoStatusReadyForDelivery: {PedidoStatusInDelivery},
	PedidoStatusInDelivery:       {PedidoStatusDeliveryFinished},
	PedidoStatusDeliveryFinished: {},
}

func (p *Pedido) UpdateStatus(newStatus PedidoStatus) error {
	if slices.Contains(validNextStatus[p.Status], newStatus) {
		p.Status = newStatus
		return nil
	}

	return fmt.Errorf("pedido em estado inválido para essa operação")
}
