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

func (ctrl *ThreadCommentController) FindAllByUserID(ctx *gin.Context) {
	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	var qs request.ThreadCommentFindAllByUserIDRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.threadCommentApplicationService.FindAllByUserID(service.ThreadCommentApplicationServiceFindByUserIDParams{
		Ctx:    ctx.Request.Context(),
		UserID: userID.(int),
		Qs:     qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadCommentController) FindById(ctx *gin.Context) {
	commentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なコメントID"})
		return
	}

	var qs request.ThreadCommentFindByIDRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.threadCommentApplicationService.FindByID(service.ThreadCommentApplicationServiceFindByIDParams{
		Ctx:       ctx.Request.Context(),
		CommentID: commentID,
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var body request.ThreadCommentCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		userID = 0
	}

	dto, err := ctrl.threadCommentApplicationService.Create(service.ThreadCommentApplicationServiceCreateParams{
		Ctx:             ctx.Request.Context(),
		UserID:          userID.(int),
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	parentCommentID, err := strconv.Atoi(ctx.Param("commentID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "無効なコメントIDです"})
		return
	}

	var body request.ThreadCommentCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := ctx.Get("userID")

	dto, err := ctrl.threadCommentApplicationService.Create(service.ThreadCommentApplicationServiceCreateParams{
		Ctx:             ctx.Request.Context(),
		ThreadID:        threadID,
		UserID:          userID.(int),
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
