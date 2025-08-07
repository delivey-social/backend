package types

import "github.com/google/uuid"

type CreatePedidoRequest struct {
	Items []CreatePedidoRequestItem `json:"itens" binding:"required,dive,required"`
}

type CreatePedidoRequestItem struct {
	ItemID   uuid.UUID `json:"item_id" binding:"required"`
	Quantity uint16    `json:"quantidade" binding:"required"`
}

type CardapioService interface {
	GetItemsByIDS(ids []uuid.UUID) ([]CardapioItem, error)
}

type CardapioItem struct {
	Id uuid.UUID
	Price uint32
}

type PedidoRepository interface {
	Create(items []PedidoRepositoryItem)
}

type PedidoRepositoryItem struct{
	ID uuid.UUID
	Quantity uint16
	PriceSnapshot uint32
}
