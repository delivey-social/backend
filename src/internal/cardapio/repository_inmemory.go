package cardapio

import (
	"sync"

	"github.com/google/uuid"
)

type InMemoryCardapioRepository struct {
	mu sync.RWMutex
}

func NewInMemoryCardapioRepository() CardapioRepository {
	return &InMemoryCardapioRepository{}
}

func (repo *InMemoryCardapioRepository) GetByID(id uuid.UUID) *Cardapio {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	return &Cardapio{}
}
