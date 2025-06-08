package services

import (
	"errors"

	"comida.app/src/pedido"
	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

type CreatePedidoCommand struct {
	UserEmail       string
	UserPhone       string
	ItemsId         uuid.UUID
	Cep             string
	Obervacao       string
	MetodoPagamento enums.MetodoPagamento
}

var (
	ErrInvalidCep     = errors.New("invalid CEP")
	ErrInvalidAddress = errors.New("invalid address")
	ErrCreatingPedido = errors.New("error creating pedido")
	ErrInvalidEmail   = errors.New("invalid email address")
	ErrInvalidPhone   = errors.New("invalid phone number")
)

func CreatePedido(cmd CreatePedidoCommand) (*pedido.Pedido, error) {
	email, err := shared.NewEmail(cmd.UserEmail)
	if err != nil {
		return nil, ErrInvalidEmail
	}

	phone, err := shared.NewTelefone(cmd.UserPhone)
	if err != nil {
		return nil, ErrInvalidPhone
	}
	usuario := shared.NewUsuario(email, phone)

	// TODO: Fetch items snapshot in restaurant context?
	var ItemsSnapshot []valueobject.ItemPedidoSnapshot
	var precoItens int

	for _, item := range ItemsSnapshot {
		precoItens += item.Preco()
	}

	cep, err := shared.NewCEP(cmd.Cep)
	if err != nil {
		return nil, ErrInvalidAddress
	}
	endereco, err := shared.NewEndereco(cep)
	if err != nil {
		return nil, ErrInvalidAddress
	}

	// TODO: Calculate delivery fee
	var taxaEntrega int

	preco := valueobject.NewPreco(precoItens, taxaEntrega)

	pedido, err := pedido.NewPedido(usuario, ItemsSnapshot, endereco, preco, cmd.Obervacao, cmd.MetodoPagamento)
	if err != nil {
		return nil, ErrCreatingPedido
	}

	// TODO: Save user in UserRepository

	return &pedido, nil
}
