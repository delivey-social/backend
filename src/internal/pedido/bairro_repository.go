package pedido

import (
	"sync"

	"github.com/google/uuid"
)

type BairroRepository interface {
	Save(bairro Bairro)
	List() []Bairro
}

type InMemoryBairroRepository struct {
	mu    sync.RWMutex
	store map[uuid.UUID]*Bairro
}

func NewInMemoryBairroRepository() BairroRepository {
	return &InMemoryBairroRepository{
		store: make(map[uuid.UUID]*Bairro),
	}
}

func (r *InMemoryBairroRepository) Save(bairro Bairro) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[bairro.ID] = &bairro
}

func (r *InMemoryBairroRepository) List() []Bairro {
	r.mu.Lock()
	defer r.mu.Unlock()

	bairros := make([]Bairro, 0, len(r.store))
	for _, bairro := range r.store {
		bairros = append(bairros, *bairro)
	}
	return bairros
}
