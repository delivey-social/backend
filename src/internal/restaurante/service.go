package restaurante

import (
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

func (s *RestauranteService) List() []Restaurante {
	return s.repo.List()
}

func (s *RestauranteService) Create(CNPJ CNPJ, Name string) uuid.UUID {
	return s.repo.Create(CNPJ, Name)
}

func (s *RestauranteService) GetMenu(restaurantID uuid.UUID) (*Cardapio, error) {
	return s.repo.GetMenu(restaurantID)
}

func (s *RestauranteService) GetMenuItemsByIDs(ids []uuid.UUID) []CardapioItem {
	return s.repo.GetItemsByIDs(ids)
}

func (s *RestauranteService) CreateMenuItem(restaurantID uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	return s.repo.CreateMenuItem(restaurantID, data)
}

func (s *RestauranteService) UpdateMenuItem(restaurantID uuid.UUID, id uuid.UUID, data MenuItemParams) error {
	s.repo.UpdateMenuItem(restaurantID, id, data)
	return nil
}

func (s *RestauranteService) DeleteMenuItem(restaurantID uuid.UUID, id uuid.UUID) error {
	s.repo.DeleteMenuItem(restaurantID, id)
	return nil
}
