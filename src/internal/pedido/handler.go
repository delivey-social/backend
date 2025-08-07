package pedido

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PedidoHandler struct{}

func NewPedidoHandler() *PedidoHandler {
	return &PedidoHandler{}
}

func (h *PedidoHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/pedido", h.create)
}

type CreatePedidoRequest struct {
	Items []CreatePedidoRequestItem `json:"itens" binding:"required,dive,required"`
}

type CreatePedidoRequestItem struct {
	ItemId uuid.UUID `json:"item_id" binding:"required"` 
	Quantity uint32 `json:"quantidade" binding:"required"`
}


func (h *PedidoHandler) create(c *gin.Context) {
	var body CreatePedidoRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "Pedido inválido",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Pedido criado com sucesso!",
	} )
}