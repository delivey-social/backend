package restaurante

import "github.com/google/uuid"

type RestauranteRepository interface {
	List() []Restaurante
}

type Restaurante struct {
	ID   uuid.UUID
	CNPJ string
	Name string
}
