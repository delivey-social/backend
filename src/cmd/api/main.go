package api

import (
	"fmt"
	"net/http"

	"comida.app/src/internal/cardapio"
	"comida.app/src/internal/pedido"
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

	router.Run(":" + PORT)
}

type InMemoryCardapioService struct{}

func (f *InMemoryCardapioService) GetItemsByIDS(ids []uuid.UUID) ([]pedido.CardapioItem, error) {
	var items = []pedido.CardapioItem{{Id: uuid.New(), Price: 1234}}

	fmt.Println(items)

	return items, nil
}
