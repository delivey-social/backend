package bairro

import (
	"fmt"

	"github.com/google/uuid"
)

type Bairro struct {
	ID          uuid.UUID `json:"id"`
	Nome        string    `json:"nome"`
	TaxaEntrega uint32    `json:"taxa_entrega"`
}

func NewBairro(nome string, taxaEntrega uint32) (Bairro, error) {
	if nome == "" || taxaEntrega == 0 {
		return Bairro{}, fmt.Errorf("invariant failed: criando bairro")
	}

	return Bairro{
		ID:          uuid.New(),
		Nome:        nome,
		TaxaEntrega: taxaEntrega,
	}, nil
}
