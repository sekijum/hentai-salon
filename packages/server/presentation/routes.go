package routes

import (
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine, controllers *di.ControllersSet) {
	{
		router.GET("/health-check", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "ok"})
		})
	}

	clientGroup := router.Group("/client")
	{
		clientGroup.POST("/boards", controllers.BoardClientController.Create)
	}
}
