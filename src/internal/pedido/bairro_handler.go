package pedido

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BairroHandler struct {
	service BairroService
}

func NewBairroHandler(service BairroService) *BairroHandler {
	return &BairroHandler{
		service,
	}
}

func (h *BairroHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/bairro", h.listBairros)
}

func (h *BairroHandler) listBairros(c *gin.Context) {
	bairros := h.service.ListBairros()

	c.JSON(http.StatusOK, gin.H{
		"bairros": bairros,
	})
}
