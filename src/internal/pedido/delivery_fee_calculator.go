package pedido

// Contains the info needed to calculate the delivery fee
type DeliveryFeeContext struct {
	bairro Bairro
}

type DeliveryFeeCalculator interface {
	Calculate(ctx DeliveryFeeContext) uint32
}

type fixedRateCalculator struct {
}

func NewFixedRateCalculator() DeliveryFeeCalculator {
	return &fixedRateCalculator{}
}

func (c *fixedRateCalculator) Calculate(ctx DeliveryFeeContext) uint32 {
	return ctx.bairro.TaxaEntrega
}
