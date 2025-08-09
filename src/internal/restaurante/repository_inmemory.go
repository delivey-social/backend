package restaurante

import (
	"sync"

	"github.com/google/uuid"
)

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

func (r *InMemoryRestauranteRepository) Create(CNPJ CNPJ, Name string) uuid.UUID {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New()
	r.store = append(r.store, Restaurante{
		ID:   id,
		CNPJ: CNPJ.String(),
		Name: Name,
		Cardapio: Cardapio{
			ID:      uuid.New(),
			Content: map[string]CardapioItem{},
		},
	})

	return id
}

func (r *InMemoryRestauranteRepository) GetItemsByIDs(ids []uuid.UUID) []CardapioItem {
	return []CardapioItem{}
}

func (r *InMemoryRestauranteRepository) GetMenu(restauranteID uuid.UUID) Cardapio {
	return Cardapio{}
}
