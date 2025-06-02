package valueobject

import "github.com/google/uuid"

type ItemPedido struct {
	ItemId     uuid.UUID
	Nome       string
	Preco      int
	Quantidade int
}
