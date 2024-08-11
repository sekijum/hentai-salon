package routes

import (
	"server/presentation/middleware"
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine, controllers *di.ControllersSet) {
	r.Use(middleware.RequestMiddleware(), middleware.OptionalAuthMiddleware())

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.GET("/health-check", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "ok"})
	})

	r.POST("/contact", controllers.ContactController.Create)

	r.POST("/signup", controllers.UserController.Signup)
	r.POST("/signin", controllers.UserController.Signin)
	r.POST("/verify-email-token", controllers.UserController.VerifyEmailToken)

	r.POST("/forgot-password", controllers.UserController.ForgotPassword)
	r.POST("/verify-reset-password-token", controllers.UserController.VerifyResetPasswordToken)
	r.PATCH("/reset-password", controllers.UserController.ResetPassword)

	r.GET("/tags/name", controllers.TagController.FindNameList)

	r.GET("/boards", controllers.BoardController.FindAll)

	r.POST("/files/urls-for-upload", controllers.StorageController.GeneratePresignedURLs)

	threadGroup := r.Group("/threads")
	{
		threadGroup.GET("", controllers.ThreadController.FindAll)
		threadGroup.GET("/:threadID", controllers.ThreadController.FindById)
		threadGroup.POST("/:threadID/like", controllers.ThreadController.Like)
		threadGroup.POST("/:threadID/unlike", controllers.ThreadController.Unlike)

		commentGroup := threadGroup.Group("/:threadID/comments")
		{
			commentGroup.GET("/:commentID", controllers.ThreadCommentController.FindById)
			commentGroup.POST("", controllers.ThreadCommentController.Create)
			commentGroup.POST("/:commentID/reply", controllers.ThreadCommentController.Reply)
			commentGroup.POST("/:commentID/like", controllers.ThreadCommentController.Like)
			commentGroup.POST("/:commentID/unlike", controllers.ThreadCommentController.Unlike)
		}
	}

	// 認証必須ルート
	authGroup := r.Group("/")
	authGroup.Use(middleware.AuthMiddleware())
	{

		threadGroup := authGroup.Group("/threads")
		{
			threadGroup.POST("", controllers.ThreadController.Create)
			threadGroup.PUT("/:threadID", controllers.ThreadController.Update)
		}

		userGroup := authGroup.Group("/users")
		{
			userGroup.GET("/me", controllers.UserController.FindAuthenticatedUser)
			userGroup.GET("/me/comments", controllers.UserController.FindComments)
			userGroup.GET("/me/threads", controllers.UserController.FindThreads)
			userGroup.GET("/me/liked-threads", controllers.UserController.FindLikedThreads)
			userGroup.GET("/me/liked-comments", controllers.UserController.FindLikedComments)
			userGroup.PUT("/me", controllers.UserController.Update)
			userGroup.PATCH("/me/password", controllers.UserController.UpdatePassword)
		}
	}

	// 管理者ルート
	adminGroup := authGroup.Group("/admin")
	adminGroup.Use(middleware.AdminMiddleware())
	{
		adminGroup.POST("/boards", controllers.BoardController.Create)

		userGroup := adminGroup.Group("/users")
		{
			userGroup.GET("", controllers.UserAdminController.FindAll)
			userGroup.GET("/:userID", controllers.UserAdminController.FindByID)
			userGroup.PATCH("/:userID/status", controllers.UserAdminController.UpdateStatus)
			userGroup.PUT("/:userID", controllers.UserAdminController.Update)
		}

		contactGroup := adminGroup.Group("/contacts")
		{
			contactGroup.GET("", controllers.ContactAdminController.FindAll)
			contactGroup.GET("/:contactID", controllers.ContactAdminController.FindByID)
			contactGroup.PATCH("/:contactID/status", controllers.ContactAdminController.UpdateStatus)
		}

		boardGroup := adminGroup.Group("/boards")
		{
			boardGroup.GET("", controllers.BoardAdminController.FindAll)
			boardGroup.GET("/:boardID", controllers.BoardAdminController.FindByID)
			boardGroup.PUT("/:boardID", controllers.BoardAdminController.Update)
			boardGroup.PATCH("/:boardID/status", controllers.BoardAdminController.UpdateStatus)
		}

		tagGroup := adminGroup.Group("/tags")
		{
			tagGroup.GET("", controllers.TagAdminController.FindAll)
			tagGroup.DELETE("/:tagID", controllers.TagAdminController.Delete)
		}

		threadGroup := adminGroup.Group("/threads")
		{
			threadGroup.GET("", controllers.ThreadAdminController.FindAll)
			threadGroup.GET("/:threadID", controllers.ThreadAdminController.FindByID)
			threadGroup.PUT("/:threadID", controllers.ThreadAdminController.Update)
			threadGroup.PATCH("/:threadID/status", controllers.ThreadAdminController.UpdateStatus)

			commentGroup := threadGroup.Group("/:threadID/comments")
			{
				commentGroup.DELETE("/:commentID", controllers.ThreadCommentAdminController.Delete)
			}
		}
	}

}
