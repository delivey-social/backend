package api

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Handlers interface {
	RegisterRoutes(r *gin.Engine)
}

func Start(handlers []Handlers) {
	PORT := os.Getenv("PORT")
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is online!",
		})
	})

	for _, handler := range handlers {
		handler.RegisterRoutes(router)
	}

	router.Run(":" + PORT)
}
