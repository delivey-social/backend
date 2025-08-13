package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
	Create(CNPJ CNPJ, Name string) uuid.UUID
	GetMenu(restaurantID uuid.UUID) (*Cardapio, error)

	GetItemsByIDs(restaurantID uuid.UUID, ids []uuid.UUID) (*[]CardapioItem, error)

	CreateMenuItem(restaurantID uuid.UUID, data MenuItemParams) (uuid.UUID, error)
	UpdateMenuItem(restaurantID uuid.UUID, ID uuid.UUID, data MenuItemParams) error
	DeleteMenuItem(restaurantID uuid.UUID, ID uuid.UUID) error
}

type Restaurante struct {
	ID       uuid.UUID `json:"id"`
	CNPJ     string    `json:"cnpj"`
	Name     string    `json:"nome"`
	Cardapio Cardapio  `json:"cardapio"`
}

type Cardapio []CardapioItem

type CardapioItem struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"nome"`
	Price    uint32    `json:"preco"`
	Category string    `json:"categoria"`
}

type MenuItemParams struct {
	Name     string
	Price    uint32
	Category string
}
