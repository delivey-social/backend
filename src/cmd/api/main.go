package api

import (
	"fmt"
	"net/http"

	"comida.app/src/internal/pedido/application"
	"comida.app/src/internal/pedido/handler"
	"comida.app/src/internal/pedido/repositories"
	"comida.app/src/internal/pedido/types"
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

	pedidoRepo := repositories.NewInMemoryPedidoRepository()

	pedidoService := application.NewPedidoService(pedidoRepo, &InMemoryCardapioService{})

	pedidoHandler := handler.NewPedidoHandler(*pedidoService)
	pedidoHandler.RegisterRoutes(router)

	router.Run(":" + PORT)
}

type InMemoryCardapioService struct{}

func (f *InMemoryCardapioService) GetItemsByIDS(ids []uuid.UUID) ([]types.CardapioItem, error) {
	var items = []types.CardapioItem{{Id: uuid.New(), Price: 1234}}

	fmt.Println(items)

	return items, nil
}
