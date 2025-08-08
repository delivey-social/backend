package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
	Create(CNPJ CNPJ, Name string, CardapioID uuid.UUID) uuid.UUID
}

type Restaurante struct {
	ID         uuid.UUID `json:"id"`
	CNPJ       string    `json:"cnpj"`
	Name       string    `json:"nome"`
	CardapioID uuid.UUID `json:"cardapio_id"`
}
