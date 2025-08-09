package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
	Create(CNPJ CNPJ, Name string) uuid.UUID
	GetMenu(restaurantID uuid.UUID) Cardapio
	GetItemsByIDs(ids []uuid.UUID) []CardapioItem
}

type Restaurante struct {
	ID       uuid.UUID `json:"id"`
	CNPJ     string    `json:"cnpj"`
	Name     string    `json:"nome"`
	Cardapio Cardapio  `json:"cardapio"`
}

type Cardapio struct {
	ID      uuid.UUID
	Content map[string]CardapioItem
}

type CardapioItem struct {
	ID    uuid.UUID
	Name  string
	Price uint32
}
