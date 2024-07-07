package controller

import (
	"context"
	"errors"
	"net/http"
	"server/application/service"
	request "server/presentation/request"

	"strconv"

	"github.com/gin-gonic/gin"
)

var ErrNotFound = errors.New("not found")

type ThreadCommentController struct {
	threadCommentApplicationService *service.ThreadCommentApplicationService
}

func NewThreadCommentController(threadCommentApplicationService *service.ThreadCommentApplicationService) *ThreadCommentController {
	return &ThreadCommentController{threadCommentApplicationService: threadCommentApplicationService}
}

func (ctrl *ThreadCommentController) FindAll(c *gin.Context) {
	var qs request.ThreadCommentFindAllRequest

	if err := c.ShouldBindQuery(&qs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comments, err := ctrl.threadCommentApplicationService.FindAll(context.Background(), qs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "コメントの取得に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (ctrl *ThreadCommentController) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なコメントID"})
		return
	}

	comment, err := ctrl.threadCommentApplicationService.FindById(context.Background(), id)
	if err != nil {
		if err == ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "コメントが見つかりません"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "コメントの取得に失敗しました: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (ctrl *ThreadCommentController) Create(ginCtx *gin.Context) {
	var body request.ThreadCommentCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.threadCommentApplicationService.Create(context.Background(), ginCtx, body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, nil)
}

func (ctrl *ThreadCommentController) Reply(ginCtx *gin.Context) {
	var body request.ThreadCommentCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.ParentCommentId == nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "ParentCommentIdが必要です"})
		return
	}

	err := ctrl.threadCommentApplicationService.Create(context.Background(), ginCtx, body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, nil)
}
