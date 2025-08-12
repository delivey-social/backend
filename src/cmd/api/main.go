package api

import (
	"net/http"

	"comida.app/src/adapters"
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

	restauranteRepo := restaurante.NewInMemoryRestauranteRepository()
	restauranteService := restaurante.NewRestauranteService(restauranteRepo)
	restauranteHandler := restaurante.NewRestaurantHandler(*restauranteService)
	restauranteHandler.RegisterRoutes(router)

	pedidoRepo := pedido.NewInMemoryPedidoRepository()
	pedidoService := pedido.NewPedidoService(pedidoRepo, adapters.NewCardapioPedidoAdapter(restauranteService))
	pedidoHandler := pedido.NewPedidoHandler(*pedidoService)
	pedidoHandler.RegisterRoutes(router)

	initializeRestaurante(restauranteService)

	router.Run(":" + PORT)
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
