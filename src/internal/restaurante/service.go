package restaurante

import (
	"fmt"

	"github.com/google/uuid"
)

type RestauranteService struct {
	repo RestauranteRepository
}

func NewRestauranteService(repo RestauranteRepository) *RestauranteService {
	return &RestauranteService{
		repo,
	}
}

func (s *RestauranteService) List() ([]Restaurante, error) {
	restaurantes, err := s.repo.List()
	if err != nil {
		return nil, fmt.Errorf("service: failed to fetch restaurantes: %w", err)
	}

	return restaurantes, nil
}

func (s *RestauranteService) Create(CNPJ CNPJ, Name string) (uuid.UUID, error) {
	id, err := s.repo.Create(CNPJ, Name)
	if err != nil {
		return id, fmt.Errorf("service: failed to create restaurante: %w", err)
	}

	return id, nil
}

func (s *RestauranteService) GetMenu(restaurantID uuid.UUID) (*Cardapio, error) {
	return s.repo.GetMenu(restaurantID)
}

func (s *RestauranteService) GetMenuItemsByIDs(restaurantID uuid.UUID, ids []uuid.UUID) (*[]CardapioItem, error) {
	return s.repo.GetItemsByIDs(restaurantID, ids)
}

func (s *RestauranteService) CreateMenuItem(restaurantID uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	return s.repo.CreateMenuItem(restaurantID, data)
}

func (s *RestauranteService) UpdateMenuItem(restaurantID uuid.UUID, id uuid.UUID, data MenuItemParams) error {
	return s.repo.UpdateMenuItem(restaurantID, id, data)
}

func (s *RestauranteService) DeleteMenuItem(restaurantID uuid.UUID, id uuid.UUID) error {
	return s.repo.DeleteMenuItem(restaurantID, id)
}
