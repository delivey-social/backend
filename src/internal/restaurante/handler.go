package restaurante

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	router.GET("/restaurante/:id/menu", h.getMenu)
}

func (h *RestauranteHandler) list(c *gin.Context) {
	restaurantes := h.service.List()

	c.JSON(http.StatusOK, gin.H{
		"restaurantes": restaurantes,
	})
}

func (h *RestauranteHandler) getMenu(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id",
		})
		return
	}

	menu, err := h.service.GetMenu(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Repository error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"menu": menu,
	})
}
