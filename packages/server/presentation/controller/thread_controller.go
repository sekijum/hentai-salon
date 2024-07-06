package controller

import (
	"context"
	"net/http"
	"server/application/service"
	request "server/presentation/request"

	"github.com/gin-gonic/gin"
)

type ThreadController struct {
	threadApplicationService *service.ThreadApplicationService
}

func NewThreadController(threadApplicationService *service.ThreadApplicationService) *ThreadController {
	return &ThreadController{threadApplicationService: threadApplicationService}
}

func (ctrl *ThreadController) FindAll(c *gin.Context) {
	var qs request.ThreadFindAllRequest

	if err := c.ShouldBindQuery(&qs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	threads, err := ctrl.threadApplicationService.FindAll(context.Background(), qs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "スレッドの取得に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (ctrl *ThreadController) Create(ginCtx *gin.Context) {
	var body request.ThreadCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.threadApplicationService.Create(context.Background(), ginCtx, body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, nil)
}
