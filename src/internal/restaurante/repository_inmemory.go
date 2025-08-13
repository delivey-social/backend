package restaurante

import (
	"errors"
	"fmt"
	"sync"

	"comida.app/src/utils"
	"github.com/google/uuid"
)

var (
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

func (r *InMemoryRestauranteRepository) GetItemsByIDs(restaurantID uuid.UUID, ids []uuid.UUID) (*[]CardapioItem, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	restaurant, err := r.findRestaurantById(restaurantID)
	if err != nil {
		return nil, err
	}

	var result []CardapioItem
	for _, id := range ids {
		for _, item := range restaurant.Cardapio {
			if item.ID == id {
				result = append(result, item)
			}
		}
	}

	if len(result) != len(ids) {
		return nil, utils.NewResourceNotFoundError("menu item")
	}

	return &result, nil
}

func (r *InMemoryRestauranteRepository) GetMenu(restaurantID uuid.UUID) (*Cardapio, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	restaurant, err := r.findRestaurantById(restaurantID)
	if err != nil {
		return nil, err
	}

	return &restaurant.Cardapio, nil
}

func (r *InMemoryRestauranteRepository) CreateMenuItem(restaurantID uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	restaurant, err := r.findRestaurantById(restaurantID)
	if err != nil {
		return uuid.UUID{}, err
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

	item, err := r.findItemById(restaurantID, ID)
	if item == nil {
		return err
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
	r.mu.Lock()
	defer r.mu.Unlock()

	restaurant, err := r.findRestaurantById(restaurantID)
	if err != nil {
		return err
	}

	for i, item := range restaurant.Cardapio {
		if item.ID == ID {
			restaurant.Cardapio = append(restaurant.Cardapio[:i], restaurant.Cardapio[i+1:]...)
			return nil
		}
	}

	return utils.NewResourceNotFoundError("menu item")
}

func (r *InMemoryRestauranteRepository) findRestaurantById(restaurantID uuid.UUID) (*Restaurante, error) {
	for i := range r.store {
		restaurant := &r.store[i]

		fmt.Println("RESTAURANT", restaurant.Name, restaurant.ID, restaurantID)

		if restaurant.ID == restaurantID {
			return restaurant, nil
		}
	}

	return nil, utils.NewResourceNotFoundError("restaurant")
}

func (r *InMemoryRestauranteRepository) findItemById(restaurantID uuid.UUID, itemID uuid.UUID) (*CardapioItem, error) {
	restaurant, err := r.findRestaurantById(restaurantID)
	if err == nil {
		return nil, err
	}

	for i := range restaurant.Cardapio {
		item := &restaurant.Cardapio[i]
		if item.ID == itemID {
			return item, nil
		}
	}

	return nil, utils.NewResourceNotFoundError("menu item")
}
