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

func (ctrl *ThreadController) FindAllList(ginCtx *gin.Context) {
	var qs request.ThreadFindAllRequest

	if err := ginCtx.ShouldBindQuery(&qs); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ginCtx.GetInt("limit")
	qs.Offset = ginCtx.GetInt("offset")

	threads, err := ctrl.threadApplicationService.FindAllList(service.ThreadApplicationServiceFindAllListParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, threads)
}

func (ctrl *ThreadController) FindById(ginCtx *gin.Context) {
	threadID, err := strconv.Atoi(ginCtx.Param("threadID"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "無効なIDです"})
		return
	}

	var qs request.ThreadFindByIdRequest

	if err := ginCtx.ShouldBindQuery(&qs); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ginCtx.GetInt("limit")
	qs.Offset = ginCtx.GetInt("offset")

	thread, err := ctrl.threadApplicationService.FindByID(service.ThreadApplicationServiceFindByIDParams{
		Ctx:      context.Background(),
		ThreadID: threadID,
		Qs:       qs,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, thread)
}

func (ctrl *ThreadController) Create(ginCtx *gin.Context) {
	var body request.ThreadCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	threads, err := ctrl.threadApplicationService.Create(service.ThreadApplicationServiceCreateParams{
		Ctx:    context.Background(),
		GinCtx: ginCtx,
		Body:   body,
	})

	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, threads)
}
