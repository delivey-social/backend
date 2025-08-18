package pedido

import (
	"comida.app/src/infra"
	"github.com/google/uuid"
)

type EventPublisher interface {
	Publish(event infra.Event)
}

type CreatePedidoDTO struct {
	RestaurantID uuid.UUID             `json:"restaurant_id" binding:"required"`
	Items        []CreatePedidoDTOItem `json:"itens" binding:"required,dive,required"`
	User         UserDTO               `json:"usuario" binding:"required"`
	Address      AddressDTO            `json:"endereco" binding:"required"`
}

type CreatePedidoDTOItem struct {
	ItemID   uuid.UUID `json:"item_id" binding:"required"`
	Quantity uint16    `json:"quantidade" binding:"required"`
}

type UserDTO struct {
	Email string `json:"email" binding:"required"`
	Phone string `json:"telefone" binding:"required"`
	Name  string `json:"nome" binding:"required"`
}

type AddressDTO struct {
	Street       string `json:"rua" binding:"required"`
	Neighborhood string `json:"bairro" binding:"required"`
	Number       string `json:"numero" binding:"required"`
	Observation  string `json:"observacao" binding:"required"`
	CEP          string `json:"CEP" binding:"required"`
}

type RestauranteService interface {
	GetItemsByIDS(restaurantID uuid.UUID, ids []uuid.UUID) ([]CardapioItem, error)
}

type CardapioItem struct {
	Id    uuid.UUID
	Price uint32
}
