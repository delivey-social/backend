package api

import (
	"net/http"

	"comida.app/src/adapters"
	"comida.app/src/internal/cardapio"
	"comida.app/src/internal/pedido"
	"comida.app/src/internal/restaurante"
	"github.com/gin-gonic/gin"
)

const PORT = "3001"

func Start() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is online",
		})
	})

	cardapioRepo := cardapio.NewInMemoryCardapioRepository()
	cardapioService := cardapio.NewCardapioService(cardapioRepo)
	cardapioHandler := cardapio.NewCardapioHandler(*cardapioService)
	cardapioHandler.RegisterRoutes(router)

	pedidoRepo := pedido.NewInMemoryPedidoRepository()
	pedidoService := pedido.NewPedidoService(pedidoRepo, adapters.NewCardapioPedidoAdapter(cardapioService))
	pedidoHandler := pedido.NewPedidoHandler(*pedidoService)
	pedidoHandler.RegisterRoutes(router)

	restauranteRepo := restaurante.NewInMemoryRestauranteRepository()
	restauranteService := restaurante.NewRestauranteService(restauranteRepo, cardapioService)
	restauranteHandler := restaurante.NewRestaurantHandler(*restauranteService)
	restauranteHandler.RegisterRoutes(router)

	initializeRestaurante(restauranteService)

	router.Run(":" + PORT)
}

func initializeRestaurante(s *restaurante.RestauranteService) {
	cnpj, err := restaurante.NewCNPJ("18781203/0001-28")
	if err != nil {
		panic(err)
	}

	s.Create(cnpj, "Santo Crepe")
}
