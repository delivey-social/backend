package valueobject

import "github.com/google/uuid"

type ItemPedidoSnapshot struct {
	itemId     uuid.UUID
	nome       string
	preco      int
	quantidade int
}

func NewItemPedidoSnapshot(
	itemId uuid.UUID,
	nome string,
	preco int,
	quantidade int,
) ItemPedidoSnapshot {
	return ItemPedidoSnapshot{
		itemId:     itemId,
		nome:       nome,
		preco:      preco,
		quantidade: quantidade,
	}
}

func (item *ItemPedidoSnapshot) ItemId() uuid.UUID {
	return item.itemId
}
func (item *ItemPedidoSnapshot) Nome() string {
	return item.nome
}
func (item *ItemPedidoSnapshot) Preco() int {
	return item.preco
}
func (item *ItemPedidoSnapshot) Quantidade() int {
	return item.quantidade
}
