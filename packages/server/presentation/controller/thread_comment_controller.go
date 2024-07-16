package controller

import (
	"context"
	"net/http"
	"server/application/service"
	request "server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadCommentController struct {
	threadCommentApplicationService *service.ThreadCommentApplicationService
}

func NewThreadCommentController(threadCommentApplicationService *service.ThreadCommentApplicationService) *ThreadCommentController {
	return &ThreadCommentController{threadCommentApplicationService: threadCommentApplicationService}
}

func (ctrl *ThreadCommentController) FindById(ginCtx *gin.Context) {
	commentID, err := strconv.Atoi(ginCtx.Param("commentID"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "無効なコメントID"})
		return
	}

	var qs request.ThreadFindByIdRequest

	if err := ginCtx.ShouldBindQuery(&qs); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ginCtx.GetInt("limit")
	qs.Offset = ginCtx.GetInt("offset")

	comment, err := ctrl.threadCommentApplicationService.FindByID(service.ThreadCommentApplicationServiceFindByIDParams{
		Ctx:       context.Background(),
		CommentID: commentID,
		Qs:        qs,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, comment)
}

func (ctrl *ThreadCommentController) Create(ginCtx *gin.Context) {
	threadID, err := strconv.Atoi(ginCtx.Param("threadID"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var body request.ThreadCommentCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource, err := ctrl.threadCommentApplicationService.Create(service.ThreadCommentApplicationServiceCreateParams{
		Ctx:             context.Background(),
		GinCtx:          ginCtx,
		ThreadID:        threadID,
		ParentCommentID: nil,
		Body:            body,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, resource)
}

func (ctrl *ThreadCommentController) Reply(ginCtx *gin.Context) {
	threadID, err := strconv.Atoi(ginCtx.Param("threadID"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	parentCommentID, err := strconv.Atoi(ginCtx.Param("commentID"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "無効なコメントIDです"})
		return
	}

	var body request.ThreadCommentCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// parentCommentIdをポインタとして渡す
	resource, err := ctrl.threadCommentApplicationService.Create(service.ThreadCommentApplicationServiceCreateParams{
		Ctx:             context.Background(),
		GinCtx:          ginCtx,
		ThreadID:        threadID,
		ParentCommentID: &parentCommentID,
		Body:            body,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, resource)
}
