package api

import (
	"fmt"
	"net/http"

	"comida.app/src/internal/cardapio"
	"comida.app/src/internal/pedido"
	"comida.app/src/internal/restaurante"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const PORT = "3001"

func Start() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is online",
		})
	})

	pedidoRepo := pedido.NewInMemoryPedidoRepository()
	pedidoService := pedido.NewPedidoService(pedidoRepo, &InMemoryCardapioService{})
	pedidoHandler := pedido.NewPedidoHandler(*pedidoService)
	pedidoHandler.RegisterRoutes(router)

	cardapioRepo := cardapio.NewInMemoryCardapioRepository()
	cardapioService := cardapio.NewCardapioService(cardapioRepo)
	cardapioHandler := cardapio.NewCardapioHandler(*cardapioService)
	cardapioHandler.RegisterRoutes(router)

	restauranteRepo := restaurante.NewInMemoryRestauranteRepository()
	restauranteService := restaurante.NewRestauranteService(restauranteRepo, &InMemoryCardapioService{})
	restauranteHandler := restaurante.NewRestaurantHandler(*restauranteService)
	restauranteHandler.RegisterRoutes(router)

	initializeRestaurante(restauranteService)

	router.Run(":" + PORT)
}

type InMemoryCardapioService struct{}

func (f *InMemoryCardapioService) GetItemsByIDS(ids []uuid.UUID) ([]pedido.CardapioItem, error) {
	var items = []pedido.CardapioItem{{Id: uuid.New(), Price: 1234}}

	fmt.Println(items)

	return items, nil
}

func (f *InMemoryCardapioService) Create() uuid.UUID {
	id := uuid.New()

	return id
}

func initializeRestaurante(s *restaurante.RestauranteService) {
	cnpj, err := restaurante.NewCNPJ("18781203/0001-28")
	if err != nil {
		panic(err)
	}

	s.Create(cnpj, "Santo Crepe")
}
