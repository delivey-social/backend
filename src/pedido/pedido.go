package pedido

import (
	"errors"

	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

type Pedido struct {
	id               uuid.UUID
	itens            []*valueobject.ItemPedidoSnapshot
	cliente          *shared.Usuario
	endereco         *shared.Endereco
	preco            *valueobject.Preco
	observacao       string
	status           enums.PedidoStatus
	metodo_pagamento enums.MetodoPagamento
}

var (
	ErrInvalidAddress = errors.New("invalid address")
)

func NewPedido(
	cliente shared.Usuario,
	itens []uuid.UUID,
	cep string,
	observacao string,
	metodo_pagamento enums.MetodoPagamento,
) (Pedido, error) {
	preco := valueobject.NewPreco(0, 0)
	endereco, err := shared.NewEndereco(cep)
	if err != nil {
		return Pedido{}, ErrInvalidAddress
	}

	return Pedido{
		id:               uuid.New(),
		itens:            make([]*valueobject.ItemPedidoSnapshot, 0),
		cliente:          &cliente,
		endereco:         &endereco,
		preco:            &preco,
		observacao:       observacao,
		status:           enums.StatusAguardandoPagamento,
		metodo_pagamento: metodo_pagamento,
	}, nil
}
