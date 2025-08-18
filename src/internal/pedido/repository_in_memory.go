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

func (r *InMemoryPedidoRepository) Create(
	items []PedidoItem,
	usuario Usuario,
	endereco Endereco,
	metodoPagamento PaymentMethods,
) uuid.UUID {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := uuid.New()
	r.store = append(r.store, Pedido{
		ID:            id,
		Items:         items,
		Status:        PedidoStatusCreated,
		Customer:      usuario,
		Address:       endereco,
		PaymentMethod: metodoPagamento,
	})

	return id
}

func (r *InMemoryPedidoRepository) FindByID(id uuid.UUID) (Pedido, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, pedido := range r.store {
		if pedido.ID == id {
			return pedido, nil
		}
	}

	return Pedido{}, utils.NewResourceNotFoundError("pedido")
}

func (r *InMemoryPedidoRepository) UpdateStatus(id uuid.UUID, status PedidoStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range r.store {
		if r.store[i].ID == id {
			r.store[i].Status = status
			return nil
		}
	}

	return utils.NewResourceNotFoundError("pedido")
}
