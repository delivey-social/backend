package pedido

import (
	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

type Pedido struct {
	id               uuid.UUID
	itens            []valueobject.ItemPedidoSnapshot
	cliente          *shared.Usuario
	endereco         *shared.Endereco
	preco            *valueobject.Preco
	observacao       string
	status           enums.PedidoStatus
	metodo_pagamento enums.MetodoPagamento
}

func NewPedido(
	cliente shared.Usuario,
	itens []valueobject.ItemPedidoSnapshot,
	endereco shared.Endereco,
	preco valueobject.Preco,
	observacao string,
	metodo_pagamento enums.MetodoPagamento,
) (Pedido, error) {
	return Pedido{
		id:               uuid.New(),
		itens:            itens,
		cliente:          &cliente,
		endereco:         &endereco,
		preco:            &preco,
		observacao:       observacao,
		status:           enums.StatusAguardandoPagamento,
		metodo_pagamento: metodo_pagamento,
	}, nil
}
