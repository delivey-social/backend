package pedido

import (
	"fmt"
	"sort"
	"sync"

	"github.com/google/uuid"
)

type BairroRepository interface {
	Save(bairro Bairro)
	List() []Bairro

	FindByID(id uuid.UUID) (*Bairro, error)
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

	sort.Slice(bairros, func(i, j int) bool {
		return bairros[i].Nome < bairros[j].Nome
	})

	return bairros
}

func (r *InMemoryBairroRepository) FindByID(id uuid.UUID) (*Bairro, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	bairro := r.store[id]
	if bairro == nil {
		return nil, fmt.Errorf("bairro não encontrado")
	}

	return bairro, nil
}
