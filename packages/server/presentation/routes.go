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

	commentGroup := r.Group("/comments")
	commentGroup.GET("/comments", controllers.ThreadCommentController.FindAll)
	commentGroup.GET("/:id", controllers.ThreadCommentController.FindById)
	commentGroup.POST("/comments", controllers.ThreadCommentController.Create)
	commentGroup.POST("/reply", controllers.ThreadCommentController.Reply)

	// 以下認証必須ルート
	authMiddleware := middleware.AuthMiddleware()
	paginationDefaultsMiddleware := middleware.PaginationDefaultsMiddleware()

	authGroup := r.Group("/")
	authGroup.Use(authMiddleware, paginationDefaultsMiddleware)
	{
		authGroup.POST("/threads", controllers.ThreadController.Create)
		authGroup.GET("/whoami", controllers.UserController.FindAuthenticatedUser)
	}

	adminGroup := r.Group("/admin")
	adminGroup.Use(authMiddleware, paginationDefaultsMiddleware)
	{
		adminGroup.POST("/boards", controllers.BoardAdminController.Create)
	}
}
