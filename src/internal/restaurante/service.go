package restaurante

import (
	"errors"

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

func (s *RestauranteService) CreateMenuItem() error {
	return errors.New("bot implemented")
}

func (s *RestauranteService) UpdateMenuItem() error {
	return errors.New("not implemented")
}

func (s *RestauranteService) DeleteMenuItem() error {
	return errors.New("not implemented")
}
