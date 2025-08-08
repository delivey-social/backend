package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
	Create(CNPJ CNPJ, Name string, CardapioID uuid.UUID) uuid.UUID
}

type Restaurante struct {
	ID         uuid.UUID
	CNPJ       string
	Name       string
	CardapioID uuid.UUID
}
