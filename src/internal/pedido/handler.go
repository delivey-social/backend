package pedido

import (
	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	service PedidoService
}

func NewPedidoHandler(service PedidoService) *PedidoHandler {
	return &PedidoHandler{
		service,
	}
}

func (h *PedidoHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/pedido", h.create)
}
