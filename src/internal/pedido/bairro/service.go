package bairro

import "github.com/google/uuid"

type BairroService struct {
	repo BairroRepository
}

func NewBairroService(repository BairroRepository) BairroService {
	service := BairroService{
		repo: repository,
	}

	PopulateBairros(&service)

	return service
}

func (s *BairroService) CreateBairro(nome string, taxaEntrega uint32) {
	bairro := NewBairro(nome, taxaEntrega)

	s.repo.Save(bairro)
}

func (s *BairroService) ListBairros() []Bairro {
	return s.repo.List()
}

func (s *BairroService) FindByID(id uuid.UUID) (*Bairro, error) {
	return s.repo.FindByID(id)
}
