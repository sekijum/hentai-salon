package controller

import (
	"context"
	"net/http"
	applicationService "server/application/service"
	request "server/presentation/request"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userApplicationService *applicationService.UserApplicationService
}

func NewUserController(userApplicationService *applicationService.UserApplicationService) *UserController {
	return &UserController{userApplicationService: userApplicationService}
}

func (ctrl *UserController) Signup(ginCtx *gin.Context) {
	var body request.UserSignupRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの内容が不正です: " + err.Error()})
		return
	}

	token, err := ctrl.userApplicationService.Signup(context.Background(), body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": "サインアップに失敗しました: " + err.Error()})
		return
	}

	ginCtx.Header("Authorization", "Bearer "+token)
	ginCtx.JSON(http.StatusOK, gin.H{"message": "サインアップが成功しました"})
}

func (ctrl *UserController) Signin(ginCtx *gin.Context) {
	var body request.UserSigninRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "リクエストの内容が不正です: " + err.Error()})
		return
	}

	token, err := ctrl.userApplicationService.Signin(context.Background(), body.Email, body.Password)
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "ログインに失敗しました: " + err.Error()})
		return
	}

	ginCtx.Header("Authorization", "Bearer "+token)
	ginCtx.JSON(http.StatusOK, gin.H{"message": "サインインが成功しました"})
}

func (ctrl *UserController) FindAuthenticatedUser(ginCtx *gin.Context) {
	authHeader := ginCtx.GetHeader("Authorization")
	if authHeader == "" {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "トークンが必要です"})
		return
	}

	// Bearer プレフィックスを取り除く
	token := strings.TrimPrefix(authHeader, "Bearer ")

	user, err := ctrl.userApplicationService.GetAuthenticatedUser(context.Background(), token)
	if err != nil {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "認証に失敗しました: " + err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, user)
}
