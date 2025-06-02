package pedido

import (
	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

type Pedido struct {
	id               uuid.UUID
	itens            []*valueobject.ItemPedido
	cliente          *shared.Usuario
	endereco         *shared.Endereco
	preco            *valueobject.Preco
	observacao       string
	status           enums.PedidoStatus
	metodo_pagamento enums.MetodoPagamento
}

func NewPedido(
	cliente shared.Usuario,
	itens []uuid.UUID,
	cep string,
	observacao string,
	metodo_pagamento enums.MetodoPagamento,
) (Pedido, error) {
	preco := valueobject.Preco{
		Preco_itens:  0,
		Taxa_app:     0,
		Taxa_entrega: 0,
	}
	endereco := shared.Endereco{
		CEP:         cep,
		Rua:         "",
		Bairro:      "",
		Numero:      "",
		Complemento: "",
		Cidade:      "",
		UF:          "",
		Latitude:    0,
		Longitude:   0,
	}

	return Pedido{
		id:               uuid.New(),
		itens:            make([]*valueobject.ItemPedido, 0),
		cliente:          &cliente,
		endereco:         &endereco,
		preco:            &preco,
		observacao:       observacao,
		status:           enums.StatusAguardandoPagamento,
		metodo_pagamento: metodo_pagamento,
	}, nil
}
