package api

import (
	"net/http"

	"comida.app/src/internal/pedido/handler"
	"github.com/gin-gonic/gin"
)

const PORT = "3001"

func Start() {
	router := gin.Default()

	router.GET("/", func (c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is online",
		})
	})

	pedidoHandler := handler.NewPedidoHandler()
	pedidoHandler.RegisterRoutes(router)

	router.Run(":" + PORT)
}