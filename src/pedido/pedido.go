package pedido

import (
	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/valueobject"
)

type Pedido struct {
	cliente          valueobject.Usuario
	endereco         valueobject.Endereco
	preco            valueobject.Preco
	observacao       string
	status           enums.PedidoStatus
	metodo_pagamento enums.MetodoPagamento
}
