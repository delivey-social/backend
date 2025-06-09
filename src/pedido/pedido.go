package pedido

import (
	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

type Pedido struct {
	ID              uuid.UUID
	Itens           []valueobject.ItemPedidoSnapshot
	Cliente         *shared.Usuario
	Endereco        *shared.Endereco
	Preco           *valueobject.Preco
	Observacao      string
	Status          enums.PedidoStatus
	MetodoPagamento enums.MetodoPagamento
}

func NewPedido(
	cliente shared.Usuario,
	itens []valueobject.ItemPedidoSnapshot,
	endereco shared.Endereco,
	preco valueobject.Preco,
	observacao string,
	metodoPagamento enums.MetodoPagamento,
) (Pedido, error) {
	return Pedido{
		ID:              uuid.New(),
		Itens:           itens,
		Cliente:         &cliente,
		Endereco:        &endereco,
		Preco:           &preco,
		Observacao:      observacao,
		Status:          enums.StatusAguardandoPagamento,
		MetodoPagamento: metodoPagamento,
	}, nil
}
