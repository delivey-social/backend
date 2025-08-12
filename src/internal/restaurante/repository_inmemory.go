package restaurante

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

var (
	ErrNotFound   = errors.New("resource not found")
	ErrUnsuported = errors.New("unsuported operation")
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
		ID:       id,
		CNPJ:     CNPJ.String(),
		Name:     Name,
		Cardapio: map[string][]CardapioItem{},
	})

	return id
}

func (r *InMemoryRestauranteRepository) GetItemsByIDs(ids []uuid.UUID) []CardapioItem {
	return []CardapioItem{}
}

func (r *InMemoryRestauranteRepository) GetMenu(restaurantID uuid.UUID) (*Cardapio, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	restaurant := r.findRestaurantById(restaurantID)
	if restaurant == nil {
		return nil, ErrNotFound
	}

	return &restaurant.Cardapio, nil
}

func (r *InMemoryRestauranteRepository) CreateMenuItem(restaurantID uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	restaurant := r.findRestaurantById(restaurantID)
	if restaurant == nil {
		return uuid.UUID{}, ErrNotFound
	}

	for category, items := range restaurant.Cardapio {
		if category == data.Category {
			id := uuid.New()

			items = append(items, CardapioItem{
				ID:       id,
				Name:     data.Name,
				Price:    data.Price,
				Category: data.Category,
			})
			restaurant.Cardapio[category] = items

			return id, nil
		}
	}

	id := uuid.New()
	restaurant.Cardapio[data.Category] = []CardapioItem{
		{
			ID:       id,
			Name:     data.Name,
			Price:    data.Price,
			Category: data.Category,
		},
	}

	return id, nil
}

func (r *InMemoryRestauranteRepository) UpdateMenuItem(id uuid.UUID, data MenuItemParams) error {
	return ErrUnsuported
}

func (r *InMemoryRestauranteRepository) DeleteMenuItem(id uuid.UUID) error {
	return ErrUnsuported
}

func (r *InMemoryRestauranteRepository) findRestaurantById(restaurantID uuid.UUID) *Restaurante {
	for _, restaurante := range r.store {
		if restaurante.ID == restaurantID {
			return &restaurante
		}
	}

	return nil
}
