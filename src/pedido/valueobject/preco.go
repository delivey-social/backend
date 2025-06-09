package valueobject

import "errors"

type Preco struct {
	preco_itens  int
	taxa_app     int
	taxa_entrega int
}

var (
	ErrInvalidItemsPrice = errors.New("items price should be greater than zero")
)

func NewPreco(preco_itens int, taxa_entrega int, taxa_app int) (Preco, error) {
	if preco_itens == 0 {
		return Preco{}, ErrInvalidItemsPrice
	}

	return Preco{
		preco_itens:  preco_itens,
		taxa_entrega: taxa_entrega,
		taxa_app:     taxa_app,
	}, nil
}

func (preco Preco) Total() int {
	return preco.preco_itens + preco.taxa_app + preco.taxa_entrega
}
