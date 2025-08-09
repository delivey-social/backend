package cardapio

import "github.com/google/uuid"

type CardapioRepository interface {
	GetByID(id uuid.UUID) *Cardapio
	Create() uuid.UUID
}

type Cardapio struct {
	ID       uuid.UUID
	Cardapio map[string]MenuItem
}

type MenuItem struct {
	ID    uuid.UUID
	Name  string
	Price uint32
}
