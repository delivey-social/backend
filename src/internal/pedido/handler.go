package pedido

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	router.POST("/pedido/:id/ready_for_delivery", h.readyForDelivery)
}

func (h *PedidoHandler) readyForDelivery(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "id inválido",
		})
		return
	}

	err = h.service.ReadyForDelivery(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "pedido atualizado",
	})
}