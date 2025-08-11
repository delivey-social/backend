package restaurante

import (
	"errors"
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
			Content: map[string][]CardapioItem{},
		},
	})

	return id
}

func (r *InMemoryRestauranteRepository) GetItemsByIDs(ids []uuid.UUID) []CardapioItem {
	return []CardapioItem{}
}

func (r *InMemoryRestauranteRepository) GetMenu(restaurantID uuid.UUID) (*Cardapio, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, restaurant := range r.store {
		if restaurant.ID == restaurantID {
			return &restaurant.Cardapio, nil
		}
	}

	return nil, errors.New("restaurant not found")
}

func (r *InMemoryRestauranteRepository) CreateMenuItem(restaurantID uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var restaurant *Restaurante
	for _, item := range r.store {
		if item.ID == restaurantID {
			restaurant = &item
			break
		}
	}

	if restaurant == nil {
		return uuid.UUID{}, errors.New("restaurant not found")
	}

	for category, items := range restaurant.Cardapio.Content {
		if category == data.Category {
			id := uuid.New()

			items = append(items, CardapioItem{
				ID:       id,
				Name:     data.Name,
				Price:    data.Price,
				Category: data.Category,
			})
			restaurant.Cardapio.Content[category] = items

			return id, nil
		}
	}

	id := uuid.New()
	restaurant.Cardapio.Content[data.Category] = []CardapioItem{
		CardapioItem{
			ID:       id,
			Name:     data.Name,
			Price:    data.Price,
			Category: data.Category,
		},
	}

	return id, nil
}

func (r *InMemoryRestauranteRepository) UpdateMenuItem(id uuid.UUID, data MenuItemParams) {
	errors.New("unsuported operation")
}

func (r *InMemoryRestauranteRepository) DeleteMenuItem(id uuid.UUID) {
	errors.New("unsuported operation")
}
