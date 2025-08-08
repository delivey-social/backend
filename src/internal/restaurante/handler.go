package restaurante

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestauranteHandler struct {
}

func NewRestaurantHandler() *RestauranteHandler {
	return &RestauranteHandler{}
}

func (h *RestauranteHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/restaurante", h.list)
}

func (h *RestauranteHandler) list(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"restaurantes": []string{},
	})
}
