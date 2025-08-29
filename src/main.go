package main

import (
	"comida.app/src/adapters"
	"comida.app/src/cmd/api"
	"comida.app/src/infra/eventbus"
	"comida.app/src/internal/notificacoes"
	"comida.app/src/internal/pedido"
	"comida.app/src/internal/restaurante"
)

func main() {
	eventBus := eventbus.NewEventBus()

	notificacoes.NewNotificacoesService(eventBus)

	restauranteRepo := restaurante.NewInMemoryRestauranteRepository()
	restauranteService := restaurante.NewRestauranteService(restauranteRepo)

	bairroRepo := pedido.NewInMemoryBairroRepository()
	bairroService := pedido.NewBairroService(bairroRepo)

	pedidoRepo := pedido.NewInMemoryPedidoRepository()
	pedidoService := pedido.NewPedidoService(
		pedidoRepo,
		adapters.NewCardapioPedidoAdapter(restauranteService),
		bairroService,
		eventBus,
	)

	initializeRestaurante(restauranteService)

	api.Start([]api.Handlers{
		restaurante.NewRestaurantHandler(*restauranteService),
		pedido.NewPedidoHandler(*pedidoService),
		pedido.NewBairroHandler(bairroService),
	})
}

func initializeRestaurante(s *restaurante.RestauranteService) {
	cnpj, err := restaurante.NewCNPJ("18781203/0001-28")
	if err != nil {
		panic(err)
	}

	id := s.Create(cnpj, "Santo Crepe")
	itemID, err := s.CreateMenuItem(id, restaurante.MenuItemParams{
		Name:     "Item de teste",
		Price:    420,
		Category: "teste",
	})

	if err != nil {
		panic(err)
	}

	s.UpdateMenuItem(id, itemID, restaurante.MenuItemParams{
		Name:     "Item de teste atualizado",
		Price:    421,
		Category: "teste",
	})

}
