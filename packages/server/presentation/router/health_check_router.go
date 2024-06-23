package router

import (
	"github.com/gin-gonic/gin"
)

func HealthCheckRouter(router *gin.Engine) {
	router.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "ok!!!!!!!"})
	})
}
