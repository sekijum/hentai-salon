package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userApplicationService *service.UserApplicationService
}

func NewUserController(userApplicationService *service.UserApplicationService) *UserController {
	return &UserController{userApplicationService: userApplicationService}
}

func (ctrl *UserController) FindByID(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	var qs request.UserFindByIdRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.userApplicationService.FindByID(service.UserApplicationServiceFindByIDParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Qs:     qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserController) Signup(ctx *gin.Context) {
	var body request.UserSignupRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.userApplicationService.Signup(service.UserApplicationServiceSignupParams{
		Ctx:  context.Background(),
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{"message": "サインアップが成功しました"})
}

func (ctrl *UserController) Signin(ctx *gin.Context) {
	var body request.UserSigninRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.userApplicationService.Signin(service.UserApplicationServiceSigninParams{
		Ctx:      context.Background(),
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if token == "" {
		ctx.JSON(http.StatusOK, gin.H{"message": "認証メールを送信しました。"})
	} else {
		ctx.Header("Authorization", "Bearer "+token)
		ctx.JSON(http.StatusOK, gin.H{"message": "サインインが成功しました"})
	}
}

func (ctrl *UserController) VerifyEmailToken(ctx *gin.Context) {
	var body request.UserVerifyEmailTokenRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := ctrl.userApplicationService.VerifyEmailToken(service.UserApplicationServiceVerifyEmailTokenParams{
		Ctx:  context.Background(),
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Authorization", "Bearer "+token)
	ctx.JSON(http.StatusOK, gin.H{"message": "サインインが成功しました"})
}

func (ctrl *UserController) FindAuthenticatedUser(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "トークンが必要です"})
		return
	}

	// Bearer プレフィックスを取り除く
	token := strings.TrimPrefix(authHeader, "Bearer ")

	user, err := ctrl.userApplicationService.GetAuthenticatedUser(service.UserApplicationGetAuthenticatedUserParams{
		Ctx:         context.Background(),
		TokenString: token,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (ctrl *UserController) Update(ctx *gin.Context) {
	var body request.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	err := ctrl.userApplicationService.Update(service.UserApplicationServiceUpdateParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Body:   body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "ユーザー情報を編集しました")
}

func (ctrl *UserController) UpdatePassword(ctx *gin.Context) {
	var body request.UserUpdatePasswordRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	err := ctrl.userApplicationService.UpdatePassword(service.UserApplicationServiceUpdatePasswordParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Body:   body,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "パスワードを編集しました")
}

func (ctrl *UserController) ForgotPassword(ctx *gin.Context) {
	var body request.UserForgotPasswordRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.userApplicationService.ForgotPassword(service.UserApplicationServiceForgotPasswordParams{
		Ctx:  context.Background(),
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "パスワード変更URLを送信しました")
}

func (ctrl *UserController) VerifyResetPasswordToken(ctx *gin.Context) {
	var body request.UserVerifyResetPasswordTokenRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.userApplicationService.VerifyResetPasswordToken(service.UserVerifyResetPasswordTokenRequestParams{
		Ctx:  context.Background(),
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "トークンが有効です")
}

func (ctrl *UserController) ResetPassword(ctx *gin.Context) {
	var body request.UserResetPasswordRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var err = ctrl.userApplicationService.ResetPassword(service.UserResetPasswordRequestParams{
		Ctx:  context.Background(),
		Body: body,
	})
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "パスワードを変更しました")
}

func (ctrl *UserController) FindLikedThreads(ctx *gin.Context) {
	var qs request.UserFindLikedCommentsRequest
	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.userApplicationService.FindLikedThreads(service.UserApplicationServiceFindLikedThreadsParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Qs:     qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserController) FindLikedComments(ctx *gin.Context) {
	var qs request.UserFindUnLikedCommentsRequest
	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.userApplicationService.FindLikedComments(service.UserApplicationServiceFindLikedCommentsParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Qs:     qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserController) FindThreads(ctx *gin.Context) {
	var qs request.UserFindThreadsRequest
	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.userApplicationService.FindThreads(service.UserApplicationServiceFindUserThreadsParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Qs:     qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserController) FindComments(ctx *gin.Context) {
	var qs request.UserFindCommentsRequest
	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	comments, err := ctrl.userApplicationService.FindUserComments(service.UserApplicationServiceFindUserCommentsParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Qs:     qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func (ctrl *UserController) Suspended(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	err := ctrl.userApplicationService.Suspended(service.UserApplicationServiceSuspendedParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
