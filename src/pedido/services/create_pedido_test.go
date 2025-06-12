package services_test

import (
	"errors"
	"testing"

	"comida.app/src/pedido"
	"comida.app/src/pedido/enums"
	"comida.app/src/pedido/services"
	"comida.app/src/pedido/valueobject"
	"comida.app/src/shared"
	"github.com/google/uuid"
)

// Mock dependencies
type mockPedidoRepository struct {
	saveErr error
}

func (m *mockPedidoRepository) Save(pedido pedido.Pedido) error {
	return m.saveErr
}

type mockRestauranteService struct {
	itemsSnapshot []valueobject.ItemPedidoSnapshot
	itemsErr      error
	coords        shared.Coordenada
	coordsErr     error
}

func (m *mockRestauranteService) GetItemsSnapshot(_ []uuid.UUID) ([]valueobject.ItemPedidoSnapshot, error) {
	return m.itemsSnapshot, m.itemsErr
}
func (m *mockRestauranteService) GetCoordinates() (shared.Coordenada, error) {
	return m.coords, m.coordsErr
}

type mockEnderecoService struct {
	endereco    shared.Endereco
	enderecoErr error
	fee         int
	feeErr      error
}

func (m *mockEnderecoService) CreateEnderecoFromCEP(_ shared.CEP) (shared.Endereco, error) {
	return m.endereco, m.enderecoErr
}
func (m *mockEnderecoService) CalculateDeliveryFee(_ shared.Coordenada, _ shared.Coordenada) (int, error) {
	return m.fee, m.feeErr
}

// Helper to create a valid command
func makeValidCommand() services.CreatePedidoCommand {
	return services.CreatePedidoCommand{
		UserEmail:       "cliente@teste.com",
		UserPhone:       "11999999999",
		ItemIds:         []uuid.UUID{uuid.New()},
		Cep:             "80000000",
		Observacao:      "Sem cebola",
		MetodoPagamento: enums.Pix,
	}
}

func makeValidItemSnapshot() valueobject.ItemPedidoSnapshot {
	item, _ := valueobject.NewItemPedidoSnapshot(uuid.New(), "Produto", 100, 1)
	return item
}

func makeValidEndereco() shared.Endereco {
	return shared.Endereco{
		Rua:        "Rua Teste",
		Numero:     "123",
		Cidade:     "Cidade Teste",
		Coordenada: shared.Coordenada{Latitude: 0, Longitude: 0},
	}
}

func TestCreatePedido_Success(t *testing.T) {
	service := services.PedidoService{
		PedidoRepository:   &mockPedidoRepository{},
		RestauranteService: &mockRestauranteService{itemsSnapshot: []valueobject.ItemPedidoSnapshot{makeValidItemSnapshot()}},
		EnderecoService:    &mockEnderecoService{endereco: makeValidEndereco(), fee: 10},
	}

	cmd := makeValidCommand()
	ped, err := service.CreatePedido(cmd)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if ped == nil {
		t.Fatal("expected pedido, got nil")
	}
	if ped.Observacao != cmd.Observacao {
		t.Errorf("expected observacao %q, got %q", cmd.Observacao, ped.Observacao)
	}
}

func TestCreatePedido_InvalidEmail(t *testing.T) {
	service := services.PedidoService{}
	cmd := makeValidCommand()
	cmd.UserEmail = "invalid-email"
	_, err := service.CreatePedido(cmd)
	if !errors.Is(err, services.ErrInvalidEmail) {
		t.Errorf("expected ErrInvalidEmail, got %v", err)
	}
}

func TestCreatePedido_InvalidPhone(t *testing.T) {
	service := services.PedidoService{}
	cmd := makeValidCommand()
	cmd.UserPhone = "invalid-phone"
	_, err := service.CreatePedido(cmd)
	if !errors.Is(err, services.ErrInvalidPhone) {
		t.Errorf("expected ErrInvalidPhone, got %v", err)
	}
}

func TestCreatePedido_FetchItemsError(t *testing.T) {
	service := services.PedidoService{
		RestauranteService: &mockRestauranteService{itemsErr: errors.New("fail")},
	}
	cmd := makeValidCommand()
	_, err := service.CreatePedido(cmd)
	if !errors.Is(err, services.ErrFetchingItems) {
		t.Errorf("expected ErrFetchingItems, got %v", err)
	}
}

func TestCreatePedido_InvalidAddress(t *testing.T) {
	service := services.PedidoService{
		RestauranteService: &mockRestauranteService{itemsSnapshot: []valueobject.ItemPedidoSnapshot{makeValidItemSnapshot()}},
		EnderecoService:    &mockEnderecoService{enderecoErr: errors.New("fail")},
	}
	cmd := makeValidCommand()
	_, err := service.CreatePedido(cmd)
	if !errors.Is(err, services.ErrInvalidAddress) {
		t.Errorf("expected ErrInvalidAddress, got %v", err)
	}
}

func TestCreatePedido_PriceError(t *testing.T) {
	service := services.PedidoService{
		RestauranteService: &mockRestauranteService{itemsSnapshot: []valueobject.ItemPedidoSnapshot{}}, // precoItens = 0
		EnderecoService:    &mockEnderecoService{endereco: makeValidEndereco(), fee: 10},
	}
	cmd := makeValidCommand()
	_, err := service.CreatePedido(cmd)
	if !errors.Is(err, services.ErrPrice) {
		t.Errorf("expected ErrPrice, got %v", err)
	}
}

func TestCreatePedido_SaveError(t *testing.T) {
	service := services.PedidoService{
		PedidoRepository:   &mockPedidoRepository{saveErr: errors.New("fail")},
		RestauranteService: &mockRestauranteService{itemsSnapshot: []valueobject.ItemPedidoSnapshot{makeValidItemSnapshot()}},
		EnderecoService:    &mockEnderecoService{endereco: makeValidEndereco(), fee: 10},
	}
	cmd := makeValidCommand()
	_, err := service.CreatePedido(cmd)
	if !errors.Is(err, services.ErrSavingPedido) {
		t.Errorf("expected ErrSavingPedido, got %v", err)
	}
}
