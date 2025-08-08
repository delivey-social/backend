package restaurante

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestaurantHandler struct {
}

func NewRestaurantHandler() *RestaurantHandler {
	return &RestaurantHandler{}
}

func (h *RestaurantHandler) RegisterRoutes(router *gin.Engine) {
	router.GET("/restaurante", h.list)
}

func (h *RestaurantHandler) list(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"restaurantes": []string{},
	})
}
