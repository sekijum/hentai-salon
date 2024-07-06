package routes

import (
	"server/presentation/middleware"
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, controllers *di.ControllersSet) {
	paginationDefaultsMiddleware := middleware.PaginationDefaultsMiddleware()

	r.Use(paginationDefaultsMiddleware)

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)

	r.GET("/tags/names", controllers.TagController.FindAllName)
	r.GET("/boards", controllers.BoardController.FindAll)
	r.GET("/threads", controllers.ThreadController.FindAll)

	commentGroup := r.Group("/comments")
		commentGroup.
		GET("/comments", controllers.ThreadCommentController.FindAll).
		GET("/:id", controllers.ThreadCommentController.FindById).
		POST("/comments", controllers.ThreadCommentController.Create).
		POST("/reply", controllers.ThreadCommentController.Reply)

	// 以下認証必須ルート
	authMiddleware := middleware.AuthMiddleware()

	authGroup := r.Group("/").Use(authMiddleware)
	{
		authGroup.POST("/threads", controllers.ThreadController.Create)
		authGroup.GET("/whoami", controllers.UserController.FindAuthenticatedUser)
	}

	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("/boards", controllers.BoardAdminController.Create)
	}
}
