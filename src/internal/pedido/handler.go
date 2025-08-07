package pedido

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PedidoHandler struct{}

func NewPedidoHandler() *PedidoHandler {
	return &PedidoHandler{}
}

func (h *PedidoHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/pedido", h.create)
}

func (h *PedidoHandler) create(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Pedido criado com sucesso!",
	} )
}