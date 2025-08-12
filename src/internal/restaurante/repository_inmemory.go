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
		Cardapio: []CardapioItem{},
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

	id := uuid.New()
	restaurant.Cardapio = append(restaurant.Cardapio, CardapioItem{
		ID:       id,
		Name:     data.Name,
		Price:    data.Price,
		Category: data.Category,
	})

	return id, nil
}

func (r *InMemoryRestauranteRepository) UpdateMenuItem(restaurantID uuid.UUID, ID uuid.UUID, data MenuItemParams) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	item := r.findItemById(restaurantID, ID)
	if item == nil {
		return ErrNotFound
	}

	*item = CardapioItem{
		ID:       item.ID,
		Name:     data.Name,
		Price:    data.Price,
		Category: data.Category,
	}

	return nil
}

func (r *InMemoryRestauranteRepository) DeleteMenuItem(restaurantID uuid.UUID, ID uuid.UUID) error {
	return ErrUnsuported
}

func (r *InMemoryRestauranteRepository) findRestaurantById(restaurantID uuid.UUID) *Restaurante {
	for i := range r.store {
		restaurant := &r.store[i]
		if restaurant.ID == restaurantID {
			return restaurant
		}
	}

	return nil
}

func (r *InMemoryRestauranteRepository) findItemById(restaurantID uuid.UUID, itemID uuid.UUID) *CardapioItem {
	restaurant := r.findRestaurantById(restaurantID)
	if restaurant == nil {
		return nil
	}

	for i := range restaurant.Cardapio {
		item := &restaurant.Cardapio[i]
		if item.ID == itemID {
			return item
		}
	}

	return nil
}
