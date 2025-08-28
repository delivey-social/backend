package pedido

import "github.com/google/uuid"

type Bairro struct {
	ID          uuid.UUID `json:"id"`
	Nome        string    `json:"nome"`
	TaxaEntrega uint32    `json:"taxa_entrega"`
}

func NewBairro(nome string, taxaEntrega uint32) Bairro {
	return Bairro{
		ID:          uuid.New(),
		Nome:        nome,
		TaxaEntrega: taxaEntrega,
	}
}
