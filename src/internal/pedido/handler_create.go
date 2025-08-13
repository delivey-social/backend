package pedido

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *PedidoHandler) create(c *gin.Context) {
	var body CreatePedidoRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "Pedido inválido",
		})
		return
	}

	err := h.service.Create(body.RestaurantID, body.Items)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Pedido criado com sucesso!",
	})
}
