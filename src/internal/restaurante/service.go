package restaurante

import "github.com/google/uuid"

type RestauranteService struct {
	repo            RestauranteRepository
	cardapioService CardapioService
}

type CardapioService interface {
	Create() uuid.UUID
}

func NewRestauranteService(repo RestauranteRepository, cardapioService CardapioService) *RestauranteService {
	return &RestauranteService{
		repo,
		cardapioService,
	}
}

func (s *RestauranteService) List() []Restaurante {
	return s.repo.List()
}

func (s *RestauranteService) Create(CNPJ CNPJ, Name string) uuid.UUID {
	cardapioID := s.cardapioService.Create()

	return s.repo.Create(CNPJ, Name, cardapioID)
}
