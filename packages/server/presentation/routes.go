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

	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)

	r.GET("/boards", controllers.BoardController.FindAll)
	r.GET("/threads", controllers.ThreadController.FindAll)

	// 以下認証必須ルート
	authMiddleware := middleware.AuthMiddleware()

	authGroup := r.Group("/")
	authGroup.Use(authMiddleware)
	{
		authGroup.POST("/threads", controllers.ThreadController.Create)
		authGroup.GET("/users/me", controllers.UserController.FindAuthenticatedUser)
	}

	adminGroup := r.Group("/admin")
	adminGroup.Use(authMiddleware)
	{
		adminGroup.POST("/boards", controllers.BoardAdminController.Create)
	}
}
