package pedido

// Contains the info needed to calculate the delivery fee
type DeliveryFeeContext struct{}

type DeliveryFeeCalculator interface {
	Calculate(ctx DeliveryFeeContext) uint32
}

type fixedRateCalculator struct {
	fee uint32
}

func NewFixedRateCalculator(fee uint32) DeliveryFeeCalculator {
	return &fixedRateCalculator{
		fee: fee,
	}
}

func (c *fixedRateCalculator) Calculate(ctx DeliveryFeeContext) uint32 {
	return c.fee
}
