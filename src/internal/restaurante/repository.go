package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
	Create(CNPJ CNPJ, Name string) uuid.UUID
	GetMenu(restaurantID uuid.UUID) (*Cardapio, error)
	GetItemsByIDs(ids []uuid.UUID) []CardapioItem

	CreateMenuItem(data MenuItemParams, restaurantId uuid.UUID) (uuid.UUID, error)
	UpdateMenuItem(id uuid.UUID, data MenuItemParams)
	DeleteMenuItem(id uuid.UUID)
}

type Restaurante struct {
	ID       uuid.UUID `json:"id"`
	CNPJ     string    `json:"cnpj"`
	Name     string    `json:"nome"`
	Cardapio Cardapio  `json:"cardapio"`
}

type Cardapio struct {
	ID      uuid.UUID
	Content map[string][]CardapioItem
}

type CardapioItem struct {
	ID       uuid.UUID
	Name     string
	Price    uint32
	Category string
}

type MenuItemParams struct {
	Name     string
	Price    uint32
	Category string
}
