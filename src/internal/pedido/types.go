package pedido

import (
	"comida.app/src/infra"
	"github.com/google/uuid"
)

type EventPublisher interface {
	Publish(event infra.Event)
}

type CreatePedidoRequest struct {
	RestaurantID uuid.UUID                 `json:"restaurant_id" binding:"required"`
	Items        []CreatePedidoRequestItem `json:"itens" binding:"required,dive,required"`
	User         UserRequest               `json:"usuario" binding:"required"`
}

type CreatePedidoRequestItem struct {
	ItemID   uuid.UUID `json:"item_id" binding:"required"`
	Quantity uint16    `json:"quantidade" binding:"required"`
}

type UserRequest struct {
	Email string `json:"email" binding:"required"`
	Phone string `json:"telefone" binding:"required"`
	Name  string `json:"nome" binding:"required"`
}

type RestauranteService interface {
	GetItemsByIDS(restaurantID uuid.UUID, ids []uuid.UUID) ([]CardapioItem, error)
}

type CardapioItem struct {
	Id    uuid.UUID
	Price uint32
}
