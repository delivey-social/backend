package restaurante

import "github.com/google/uuid"

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

func (s *RestauranteService) Create(CNPJ string, Name string) uuid.UUID {
	return s.repo.Create(CNPJ, Name)
}
