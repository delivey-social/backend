package restaurante

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestauranteHandler struct {
	service RestauranteService
}

func NewRestaurantHandler(service RestauranteService) *RestauranteHandler {
	return &RestauranteHandler{
		service,
	}
}

func (h *RestauranteHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/restaurante", h.list)
}

func (h *RestauranteHandler) list(c *gin.Context) {
	restaurantes := h.service.repo.List()

	c.JSON(http.StatusOK, gin.H{
		"restaurantes": restaurantes,
	})
}
