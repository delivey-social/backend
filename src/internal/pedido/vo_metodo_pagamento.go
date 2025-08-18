package pedido

import "errors"

var (
	ErrInvalidPaymentMethod = errors.New("método de pagamento inválido")
)

func ToMetodoPagamento(val string) (PaymentMethod, error) {
	switch PaymentMethod(val) {
	case DEBITO_RECEBIMENTO, PIX:
		return PaymentMethod(val), nil
	default:
		return "", ErrInvalidPaymentMethod
	}
}
