package handler

import (
	"github.com/gin-gonic/gin"
)

type PedidoHandler struct{}

func NewPedidoHandler() *PedidoHandler {
	return &PedidoHandler{}
}

func (h *PedidoHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/pedido", h.create)
}
