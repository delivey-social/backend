package cardapio

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CardapioHandler struct {
	service CardapioService
}

func NewCardapioHandler(service CardapioService) *CardapioHandler {
	return &CardapioHandler{
		service,
	}
}

func (h *CardapioHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/cardapio/:id", h.getDetails)
}

func (h *CardapioHandler) getDetails(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"mensagem": "id inválido",
		})
		return
	}

	h.service.GetDetails(id)

	c.JSON(http.StatusOK, gin.H{
		"cardapio": []string{},
	})
}
