package cardapio

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CardapioHandler struct{}

func NewCardapioHandler() *CardapioHandler {
	return &CardapioHandler{}
}

func (h *CardapioHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/cardapio/:id", h.getDetails)
}

func (h *CardapioHandler) getDetails(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"cardapio": []string{},
	})
}
