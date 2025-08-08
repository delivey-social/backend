package cardapio

import (
	"github.com/google/uuid"
)

type CardapioService struct {
	repository CardapioRepository
}

type CardapioRepository interface {
	GetByID(id uuid.UUID) *Cardapio
}

type Cardapio map[string]MenuItem

type MenuItem struct {
	ID    uuid.UUID
	Name  string
	Price uint32
}

func NewCardapioService() *CardapioService {
	return &CardapioService{}
}

func (s *CardapioService) GetDetails(id uuid.UUID) Cardapio {
	return *s.repository.GetByID(id)
}
