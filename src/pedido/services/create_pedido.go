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
	ItemIds         []uuid.UUID
	Cep             string
	Observacao      string
	MetodoPagamento enums.MetodoPagamento
}

var (
	ErrInvalidCep     = errors.New("invalid CEP")
	ErrInvalidAddress = errors.New("invalid address")
	ErrInvalidEmail   = errors.New("invalid email address")
	ErrInvalidPhone   = errors.New("invalid phone number")
	ErrPrice          = errors.New("invalid price")
	ErrCreatingPedido = errors.New("error creating pedido")
	ErrFetchingItems  = errors.New("error while fetching items from restaurante")
	ErrSavingPedido   = errors.New("error while saving pedido")
)

const TAXA_APP_PERCENTAGE = 0.1

func (service PedidoService) CreatePedido(cmd CreatePedidoCommand) (*pedido.Pedido, error) {
	usuario, err := service.createUsuario(cmd.UserEmail, cmd.UserPhone)
	if err != nil {
		return nil, err
	}

	itens, precoItens, err := service.getItems(cmd.ItemIds)
	if err != nil {
		return nil, err
	}
	var taxaApp int = int(float32(precoItens) * TAXA_APP_PERCENTAGE)

	endereco, err := service.getEndereco(cmd.Cep)
	if err != nil {
		return nil, err
	}

	taxaEntrega, err := service.calculateTaxaEntrega(endereco)
	if err != nil {
		return nil, err
	}

	preco, err := valueobject.NewPreco(precoItens, taxaEntrega, taxaApp)
	if err != nil {
		return nil, ErrPrice
	}

	pedido, err := pedido.NewPedido(usuario, itens, endereco, preco, cmd.Observacao, cmd.MetodoPagamento)
	if err != nil {
		return nil, ErrCreatingPedido
	}

	err = service.PedidoRepository.Save(pedido)
	if err != nil {
		return nil, ErrSavingPedido
	}

	return &pedido, nil
}

func (s PedidoService) createUsuario(inputEmail string, inputTelefone string) (shared.Usuario, error) {
	email, err := shared.NewEmail(inputEmail)
	if err != nil {
		return shared.Usuario{}, ErrInvalidEmail
	}

	phone, err := shared.NewTelefone(inputTelefone)
	if err != nil {
		return shared.Usuario{}, ErrInvalidPhone
	}

	return shared.NewUsuario(email, phone), nil
}

func (s PedidoService) getItems(ids []uuid.UUID) ([]valueobject.ItemPedidoSnapshot, int, error) {
	itemsSnapshot, err := s.RestauranteService.GetItemsSnapshot(ids)
	if err != nil {
		return nil, 0, ErrFetchingItems
	}

	var preco int
	for _, item := range itemsSnapshot {
		preco += item.Preco
	}

	return itemsSnapshot, preco, nil
}

func (s PedidoService) calculateTaxaEntrega(endereco shared.Endereco) (int, error) {
	restauranteCoordinates, err := s.RestauranteService.GetCoordinates()
	if err != nil {
		return 0, errors.New("error while getting restaurant coordinates")
	}

	return s.EnderecoService.CalculateDeliveryFee(restauranteCoordinates, endereco.Coordenada)
}

func (s PedidoService) getEndereco(cepInput string) (shared.Endereco, error) {
	cep, err := shared.NewCEP(cepInput)
	if err != nil {
		return shared.Endereco{}, ErrInvalidAddress
	}

	endereco, err := s.EnderecoService.CreateEnderecoFromCEP(cep)
	if err != nil {
		return shared.Endereco{}, ErrInvalidAddress
	}

	return endereco, nil
}
