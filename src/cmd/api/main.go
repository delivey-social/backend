package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const PORT = "3001"

type Handlers interface {
	RegisterRoutes(r *gin.Engine)
}

func Start(handlers []Handlers) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is online",
		})
	})

	for _, handler := range handlers {
		handler.RegisterRoutes(router)
	}

	router.Run(":" + PORT)
}

