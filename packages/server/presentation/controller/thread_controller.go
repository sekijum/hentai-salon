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

func (ctrl *ThreadController) FindAll(c *gin.Context) {
	var qs request.ThreadFindAllRequest

	if err := c.ShouldBindQuery(&qs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = c.GetInt("limit")
	qs.Offset = c.GetInt("offset")

	threads, err := ctrl.threadApplicationService.FindAll(context.Background(), qs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "スレッドの取得に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, threads)
}

func (ctrl *ThreadController) FindById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("threadId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var qs request.ThreadFindByIdRequest

	if err := c.ShouldBindQuery(&qs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = c.GetInt("limit")
	qs.Offset = c.GetInt("offset")

	thread, err := ctrl.threadApplicationService.FindById(context.Background(), id, qs)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "スレッドの取得に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, thread)
}

func (ctrl *ThreadController) Create(ginCtx *gin.Context) {
	var body request.ThreadCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	threads, err := ctrl.threadApplicationService.Create(context.Background(), ginCtx, body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, threads)
}
