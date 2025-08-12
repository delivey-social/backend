package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
	Create(CNPJ CNPJ, Name string) uuid.UUID
	GetMenu(restaurantID uuid.UUID) (*Cardapio, error)
	GetItemsByIDs(ids []uuid.UUID) []CardapioItem

	CreateMenuItem(restaurantId uuid.UUID, data MenuItemParams) (uuid.UUID, error)
	UpdateMenuItem(id uuid.UUID, data MenuItemParams) error
	DeleteMenuItem(id uuid.UUID) error
}

type Restaurante struct {
	ID       uuid.UUID `json:"id"`
	CNPJ     string    `json:"cnpj"`
	Name     string    `json:"nome"`
	Cardapio Cardapio  `json:"cardapio"`
}

type Cardapio struct {
	ID      uuid.UUID                 `json:"id"`
	Content map[string][]CardapioItem `json:"content"`
}

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
