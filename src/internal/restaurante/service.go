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

func (s *RestauranteService) CreateMenuItem(restaurantId uuid.UUID, data MenuItemParams) (uuid.UUID, error) {
	return s.repo.CreateMenuItem(restaurantId, data)
}

func (s *RestauranteService) UpdateMenuItem(id uuid.UUID, data MenuItemParams) error {
	s.repo.UpdateMenuItem(id, data)
	return nil
}

func (s *RestauranteService) DeleteMenuItem(id uuid.UUID) error {
	s.repo.DeleteMenuItem(id)
	return nil
}
