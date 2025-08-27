package pedido

import (
	"sync"

	"comida.app/src/utils"
	"github.com/google/uuid"
)

type InMemoryPedidoRepository struct {
	mu    sync.RWMutex
	store map[uuid.UUID]*Pedido
}

func NewInMemoryPedidoRepository() PedidoRepository {
	return &InMemoryPedidoRepository{
		store: make(map[uuid.UUID]*Pedido),
	}
}

func (r *InMemoryPedidoRepository) Save(pedido Pedido) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store[pedido.GetId()] = &pedido
}

func (r *InMemoryPedidoRepository) Update(pedido Pedido) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.store[pedido.GetId()]; !exists {
		return utils.NewResourceNotFoundError("pedido")
	}

	r.store[pedido.GetId()] = &pedido
	return nil
}

func (r *InMemoryPedidoRepository) FindByID(id uuid.UUID) (Pedido, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	pedido := r.store[id]
	if pedido == nil {
		return Pedido{}, utils.NewResourceNotFoundError("pedido")
	}

	return *pedido, nil

}
