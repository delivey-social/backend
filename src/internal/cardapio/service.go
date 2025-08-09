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

func (s *CardapioService) Create() uuid.UUID {

	return s.repository.Create()
}

func (s *CardapioService) GetItemsByIDS(ids []uuid.UUID) []MenuItem {
	return []MenuItem{}
}
