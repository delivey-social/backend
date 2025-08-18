package pedido

import "errors"

var (
	ErrInvalidPaymentMethod = errors.New("método de pagamento inválido")
)

func ToMetodoPagamento(val interface{}) (PaymentMethods, error) {
	res, ok := val.(string)
	if !ok {
		return "", ErrInvalidPaymentMethod
	}

	switch PaymentMethods(res) {
	case DEBITO_RECEBIMENTO, PIX:
		return PaymentMethods(res), nil
	default:
		return "", ErrInvalidPaymentMethod
	}
}
