package pedido

import (
	"sync"

	"github.com/google/uuid"
)

type InMemoryPedidoRepository struct {
	mu    sync.RWMutex
	store []Pedido
}

type Pedido struct {
	ID    uuid.UUID
	Items []PedidoItem
}

type PedidoItem struct {
	ID            uuid.UUID
	Quantity      uint16
	PriceSnapshot uint32
}

type PedidoRepository interface {
	Create(items []PedidoItem) uuid.UUID
}

func NewInMemoryPedidoRepository() PedidoRepository {
	return &InMemoryPedidoRepository{}
}

func (r *InMemoryPedidoRepository) Create(items []PedidoItem) uuid.UUID {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New()
	r.store = append(r.store, Pedido{
		ID:    id,
		Items: items,
	})

	return id
}
