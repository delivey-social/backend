package handler

import (
	"comida.app/src/internal/pedido/application"
	"github.com/gin-gonic/gin"
)

type PedidoHandler struct {
	service application.PedidoService
}

func NewPedidoHandler(service application.PedidoService) *PedidoHandler {
	return &PedidoHandler{
		service,
	}
}

func (h *PedidoHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/pedido", h.create)
}
