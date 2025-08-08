package cardapio

import (
	"sync"

	"github.com/google/uuid"
)

type InMemoryCardapioRepository struct {
	mu    sync.RWMutex
	store []*Cardapio
}

func NewInMemoryCardapioRepository() CardapioRepository {
	return &InMemoryCardapioRepository{}
}

func (repo *InMemoryCardapioRepository) GetByID(id uuid.UUID) *Cardapio {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return &Cardapio{}
}

func (repo *InMemoryCardapioRepository) Create(restaurantId uuid.UUID) uuid.UUID {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	id := uuid.New()
	repo.store = append(repo.store, &Cardapio{
		ID:           id,
		RestaurantID: restaurantId,
		Cardapio:     make(map[string]MenuItem),
	})

	return id
}
