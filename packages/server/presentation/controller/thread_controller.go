package controller

import (
	"context"
	"net/http"
	"server/application/service"
	request "server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadController struct {
	threadApplicationService *service.ThreadApplicationService
}

func NewThreadController(threadApplicationService *service.ThreadApplicationService) *ThreadController {
	return &ThreadController{threadApplicationService: threadApplicationService}
}

func (ctrl *ThreadController) FindAll(ctx *gin.Context) {
	var qs request.ThreadFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		userID = 0
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.threadApplicationService.FindAll(service.ThreadApplicationServiceFindAllParams{
		Ctx:    context.Background(),
		Qs:     qs,
		UserID: userID.(int),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadController) FindById(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var qs request.ThreadFindByIdRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	userID, exists := ctx.Get("userID")
	var userIDPtr *int
	if exists {
		id := userID.(int)
		userIDPtr = &id
	}

	dto, err := ctrl.threadApplicationService.FindByID(service.ThreadApplicationServiceFindByIDParams{
		Ctx:      context.Background(),
		UserID:   userIDPtr,
		ThreadID: threadID,
		Qs:       qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadController) Create(ctx *gin.Context) {
	var body request.ThreadCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	dto, err := ctrl.threadApplicationService.Create(service.ThreadApplicationServiceCreateParams{
		Ctx:      context.Background(),
		UserID:   userID.(int),
		ClientIP: ctx.ClientIP(),
		Body:     body,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadController) Like(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	err = ctrl.threadApplicationService.Like(service.ThreadApplicationServiceLikeParams{
		Ctx:      context.Background(),
		UserID:   userID.(int),
		ThreadID: threadID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "スレッドにいいねしました"})
}

func (ctrl *ThreadController) Unlike(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	err = ctrl.threadApplicationService.Unlike(service.ThreadApplicationServiceUnlikeParams{
		Ctx:      context.Background(),
		UserID:   userID.(int),
		ThreadID: threadID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "スレッドのいいねを外しました"})
}

func (ctrl *ThreadController) Update(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	var body request.ThreadUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.threadApplicationService.Update(service.ThreadApplicationServiceUpdateParams{
		Ctx:      context.Background(),
		ThreadID: threadID,
		Body:     body,
		UserID:   userID.(int),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadController) Delete(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.threadApplicationService.Delete(service.ThreadApplicationServiceDeleteParams{
		Ctx:      context.Background(),
		ThreadID: threadID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
