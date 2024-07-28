package controller

import (
	"net/http"
	"server/application/service"
	"server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadCommentController struct {
	threadCommentApplicationService *service.ThreadCommentApplicationService
}

func NewThreadCommentController(threadCommentApplicationService *service.ThreadCommentApplicationService) *ThreadCommentController {
	return &ThreadCommentController{threadCommentApplicationService: threadCommentApplicationService}
}

func (ctrl *ThreadCommentController) FindById(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var qs request.ThreadCommentFindByIDRequest

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

	dto, err := ctrl.threadCommentApplicationService.FindByID(service.ThreadCommentApplicationServiceFindByIDParams{
		Ctx:       ctx.Request.Context(),
		CommentID: commentID,
		UserID:    userIDPtr,
		Qs:        qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadCommentController) Create(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request.ThreadCommentCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	var userIDPtr *int
	if exists {
		id := userID.(int)
		userIDPtr = &id
	}

	dto, err := ctrl.threadCommentApplicationService.Create(service.ThreadCommentApplicationServiceCreateParams{
		Ctx:             ctx.Request.Context(),
		UserID:          userIDPtr,
		ClientIP:        ctx.ClientIP(),
		ThreadID:        threadID,
		ParentCommentID: nil,
		Body:            body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadCommentController) Reply(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parentCommentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request.ThreadCommentCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	var userIDPtr *int
	if exists {
		id := userID.(int)
		userIDPtr = &id
	}

	dto, err := ctrl.threadCommentApplicationService.Create(service.ThreadCommentApplicationServiceCreateParams{
		Ctx:             ctx.Request.Context(),
		ThreadID:        threadID,
		UserID:          userIDPtr,
		ClientIP:        ctx.ClientIP(),
		ParentCommentID: &parentCommentID,
		Body:            body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadCommentController) Like(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.threadCommentApplicationService.Like(service.ThreadCommentApplicationServiceLikeParams{
		Ctx:       ctx.Request.Context(),
		UserID:    userID.(int),
		CommentID: commentID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "liked"})
}

func (ctrl *ThreadCommentController) Unlike(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ctrl.threadCommentApplicationService.Unlike(service.ThreadCommentApplicationServiceUnLikeParams{
		Ctx:       ctx.Request.Context(),
		UserID:    userID.(int),
		CommentID: commentID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "unliked"})
}
