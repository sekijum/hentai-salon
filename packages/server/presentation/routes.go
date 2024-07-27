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

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/contact", controllers.ContactController.Create)

	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)

	r.POST("/forgot-password", controllers.UserController.ForgotPassword)
	r.POST("/verify-reset-password-token", controllers.UserController.VerifyResetPasswordToken)
	r.PATCH("/reset-password", controllers.UserController.ResetPassword)

	r.GET("/tags/name", controllers.TagController.FindNameList)

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

		usersGroup := authGroup.Group("/users")
		{
			usersGroup.GET("/:userID", controllers.UserController.FindByID)
			usersGroup.GET("/me", controllers.UserController.FindAuthenticatedUser)
			usersGroup.GET("/me/comments", controllers.ThreadCommentController.FindAllByUserID)
			usersGroup.GET("/me/threads", controllers.ThreadController.FindByUserID)
			usersGroup.PUT("/me", controllers.UserController.Update)
			usersGroup.PATCH("/me/password", controllers.UserController.UpdatePassword)
		}
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
