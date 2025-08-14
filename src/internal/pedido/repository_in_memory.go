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

func (r *InMemoryPedidoRepository) FindByID(id uuid.UUID) (Pedido, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, pedido := range r.store {
		if(pedido.ID == id) {
			return pedido, nil
		}
	}
	
	return Pedido{}, utils.NewResourceNotFoundError("pedido")
}