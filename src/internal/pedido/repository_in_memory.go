package pedido

import (
	"sync"

	"comida.app/src/utils"
	"github.com/google/uuid"
)

type InMemoryPedidoRepository struct {
	mu    sync.RWMutex
	store []Pedido
}

func NewInMemoryPedidoRepository() PedidoRepository {
	return &InMemoryPedidoRepository{}
}

func (r *InMemoryPedidoRepository) Create(pedido Pedido) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.store = append(r.store, pedido)
}

func (r *InMemoryPedidoRepository) Update(id uuid.UUID, pedido Pedido) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.store {
		if r.store[i].id == id {
			r.store[i] = pedido
			return nil
		}
	}

	return utils.NewResourceNotFoundError("pedido")
}

func (r *InMemoryPedidoRepository) FindByID(id uuid.UUID) (Pedido, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, pedido := range r.store {
		if pedido.id == id {
			return pedido, nil
		}
	}

	return Pedido{}, utils.NewResourceNotFoundError("pedido")
}
