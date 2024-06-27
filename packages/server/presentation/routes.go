package routes

import (
	"server/presentation/middleware"
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, controllers *di.ControllersSet) {
	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	authMiddleware := middleware.AuthMiddleware()

	r.GET("/users/me", authMiddleware, controllers.UserController.GetAuthenticatedUser)
	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)

	clientGroup := r.Group("/client")
	clientGroup.Use(authMiddleware)
	{
		clientGroup.POST("/boards", controllers.BoardClientController.Create)
	}
}
