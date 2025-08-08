package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
	Create(CNPJ string, Name string) uuid.UUID
}

type Restaurante struct {
	ID   uuid.UUID
	CNPJ string
	Name string
}
