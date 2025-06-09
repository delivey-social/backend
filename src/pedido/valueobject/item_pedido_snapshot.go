package valueobject

import (
	"errors"

	"github.com/google/uuid"
)

type ItemPedidoSnapshot struct {
	ItemId     uuid.UUID
	Nome       string
	Preco      int
	Quantidade int
}

var (
	ErrInvalidQuantity = errors.New("items quantity should be greater than zero")
)

func NewItemPedidoSnapshot(
	itemId uuid.UUID,
	nome string,
	preco int,
	quantidade int,
) (ItemPedidoSnapshot, error) {
	if preco <= 0 {
		return ItemPedidoSnapshot{}, ErrInvalidItemsPrice
	}

	if quantidade <= 0 {
		return ItemPedidoSnapshot{}, ErrInvalidQuantity
	}

	return ItemPedidoSnapshot{
		ItemId:     itemId,
		Nome:       nome,
		Preco:      preco,
		Quantidade: quantidade,
	}, nil
}
