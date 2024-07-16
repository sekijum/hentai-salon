package routes

import (
	"server/presentation/middleware"
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, controllers *di.ControllersSet) {
	paginationMiddleware := middleware.PaginationMiddleware()
	optionalAuthMiddleware := middleware.OptionalAuthMiddleware()

	r.Use(paginationMiddleware, optionalAuthMiddleware)

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)

	r.GET("/tags/names", controllers.TagController.FindAllName)
	r.GET("/boards", controllers.BoardController.FindAll)
	r.POST("/files/urls-for-upload", controllers.StorageController.GeneratePresignedURLs)

	threadGroup := r.Group("/threads")
	{
		threadGroup.GET("/", controllers.ThreadController.FindAllList)
		threadGroup.GET("/:threadID", controllers.ThreadController.FindById)

		commentGroup := threadGroup.Group("/:threadID/comments")
		{
			commentGroup.GET("/:commentID", controllers.ThreadCommentController.FindById)
			commentGroup.POST("/", controllers.ThreadCommentController.Create)
			commentGroup.POST("/:commentID/reply", controllers.ThreadCommentController.Reply)
		}
	}

	// 認証必須ルート
	authMiddleware := middleware.AuthMiddleware()

	authGroup := r.Group("/")
	authGroup.Use(authMiddleware)
	{
		authGroup.POST("/threads", controllers.ThreadController.Create)
		authGroup.GET("/whoami", controllers.UserController.FindAuthenticatedUser)
	}

	// 管理者ルート
	adminGroup := r.Group("/admin")
	adminGroup.Use(authMiddleware)
	{
		adminGroup.POST("/boards", controllers.BoardController.Create)

		usersGroup := adminGroup.Group("/users")
		{
			usersGroup.GET("/", controllers.UserAdminController.FindAll)
			usersGroup.PUT("/:userID", controllers.UserAdminController.Update)
		}

		boardsGroup := adminGroup.Group("/boards")
		{
			boardsGroup.GET("/", controllers.BoardAdminController.FindAll)
			boardsGroup.PUT("/:boardID", controllers.BoardAdminController.Update)
		}

		threadsGroup := adminGroup.Group("/threads")
		{
			threadsGroup.GET("/", controllers.ThreadAdminController.FindAll)
			threadsGroup.GET("/:threadID", controllers.ThreadAdminController.FindByID)
			threadsGroup.PUT("/:threadID", controllers.ThreadAdminController.Update)
		}
	}
}
