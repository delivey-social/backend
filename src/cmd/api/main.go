package api

import (
	"net/http"

	"comida.app/src/infra/environment"
	"github.com/gin-gonic/gin"
)

type Handlers interface {
	RegisterRoutes(r *gin.Engine)
}

func Start(handlers []Handlers) {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Service is online!",
		})
	})

	for _, handler := range handlers {
		handler.RegisterRoutes(router)
	}

	router.Run(":" + environment.PORT)
}
