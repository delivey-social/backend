package cardapio

import (
	"github.com/google/uuid"
)

type CardapioService struct {
	repository CardapioRepository
}

func NewCardapioService(repository CardapioRepository) *CardapioService {
	return &CardapioService{
		repository,
	}
}

func (s *CardapioService) GetDetails(id uuid.UUID) Cardapio {
	return *s.repository.GetByID(id)
}
