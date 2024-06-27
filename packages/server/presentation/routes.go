package routes

import (
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, controllers *di.ControllersSet) {
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.GET("/users/me", controllers.UserController.GetAuthenticatedUser)
	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)

	clientGroup := r.Group("/client")
	clientGroup.POST("/boards", controllers.BoardClientController.Create)
}
