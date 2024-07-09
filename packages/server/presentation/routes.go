package routes

import (
	"server/presentation/middleware"
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, controllers *di.ControllersSet) {
	paginationMiddleware := middleware.PaginationMiddleware()

	r.Use(paginationMiddleware)

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)

	r.GET("/tags/names", controllers.TagController.FindAllName)
	r.GET("/boards", controllers.BoardController.FindAll)
	r.GET("/threads", controllers.ThreadController.FindAll)

	threadGroup := r.Group("/threads")
	threadGroup.GET("/:thread_id", controllers.ThreadController.FindById)

	commentGroup := threadGroup.Group("/:thread_id/comments")
	commentGroup.
		GET("/", controllers.ThreadCommentController.FindAll).
		GET("/:comment_id", controllers.ThreadCommentController.FindById).
		POST("/", controllers.ThreadCommentController.Create).
		POST("/:comment_id/reply", controllers.ThreadCommentController.Reply)

	// 認証必須ルート
	authMiddleware := middleware.AuthMiddleware()

	authGroup := r.Group("/").Use(authMiddleware)
	{
		authGroup.POST("/threads", controllers.ThreadController.Create)
		authGroup.GET("/whoami", controllers.UserController.FindAuthenticatedUser)
	}

	// 管理者ルート
	adminGroup := r.Group("/admin").Use(authMiddleware)
	adminGroup.POST("/boards", controllers.BoardController.Create)
}
