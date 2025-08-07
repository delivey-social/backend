package repositories

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
)

type InMemoryPedidoRepository struct {
	mu sync.RWMutex
}

type PedidoRepository interface {
	Create(items []PedidoRepositoryItem)
}

type PedidoRepositoryItem struct{
	ID uuid.UUID
	Quantity uint16
	PriceSnapshot uint32
}

func NewInMemoryPedidoRepository() PedidoRepository{
	return &InMemoryPedidoRepository{}
}

func (r *InMemoryPedidoRepository) Create(items []PedidoRepositoryItem) {
	r.mu.Lock()
	defer r.mu.Unlock()

	fmt.Println("Creating new pedido")
}