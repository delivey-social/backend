package restaurante

import "sync"

type InMemoryRestauranteRepository struct {
	mu    sync.RWMutex
	store []Restaurante
}

func NewInMemoryRestauranteRepository() RestauranteRepository {
	return &InMemoryRestauranteRepository{
		store: []Restaurante{},
	}
}

func (r *InMemoryRestauranteRepository) List() []Restaurante {
	r.mu.Lock()
	defer r.mu.Unlock()

	return r.store
}
