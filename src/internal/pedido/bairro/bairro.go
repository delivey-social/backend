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
	if nome == "" {
		return Bairro{}, fmt.Errorf("nome do bairro não pode ser vazio")
	}

	if taxaEntrega == 0 {
		return Bairro{}, fmt.Errorf("taxa de entrega não pode ser zero")
	}

	return Bairro{
		ID:          uuid.New(),
		Nome:        nome,
		TaxaEntrega: taxaEntrega,
	}, nil
}
