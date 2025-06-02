package valueobject

import "github.com/google/uuid"

type ItemPedido struct {
	ItemId     uuid.UUID
	Nome       string
	Descricao  string
	Preco      int
	Quantidade int
	ImagemURL  string
}
