package pedido

import "errors"

var (
	ErrInvalidPaymentMethod = errors.New("método de pagamento inválido")
)

func ToMetodoPagamento(val interface{}) (PaymentMethod, error) {
	res, ok := val.(string)
	if !ok {
		return "", ErrInvalidPaymentMethod
	}

	switch PaymentMethod(res) {
	case DEBITO_RECEBIMENTO, PIX:
		return PaymentMethod(res), nil
	default:
		return "", ErrInvalidPaymentMethod
	}
}
