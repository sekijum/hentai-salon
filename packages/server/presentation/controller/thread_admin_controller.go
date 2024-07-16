package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ThreadAdminController struct {
	threadApplicationService *service.ThreadAdminApplicationService
}

func NewThreadAdminController(threadApplicationService *service.ThreadAdminApplicationService) *ThreadAdminController {
	return &ThreadAdminController{threadApplicationService: threadApplicationService}
}

func (ctrl *ThreadAdminController) FindAll(ginCtx *gin.Context) {
	var qs request.ThreadAdminFindAllRequest

	if err := ginCtx.ShouldBindQuery(&qs); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ginCtx.GetInt("limit")
	qs.Offset = ginCtx.GetInt("offset")

	listResource, err := ctrl.threadApplicationService.FindAll(service.ThreadAdminApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, listResource)
}

func (ctrl *ThreadAdminController) FindByID(ginCtx *gin.Context) {
	threadID, err := strconv.Atoi(ginCtx.Param("threadId"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	var qs request.ThreadAdminFindByIDRequest

	if err := ginCtx.ShouldBindQuery(&qs); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ginCtx.GetInt("limit")
	qs.Offset = ginCtx.GetInt("offset")

	listResource, err := ctrl.threadApplicationService.FindByID(service.ThreadAdminApplicationServiceFindByIDParams{
		Ctx:      context.Background(),
		ThreadID: threadID,
		Qs:       qs,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, listResource)
}

func (ctrl *ThreadAdminController) Update(ginCtx *gin.Context) {
	var body request.ThreadAdminUpdateRequest

	threadID, err := strconv.Atoi(ginCtx.Param("threadId"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid thread ID"})
		return
	}

	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource, err := ctrl.threadApplicationService.Update(service.ThreadAdminApplicationServiceUpdateParams{
		Ctx:      context.Background(),
		ThreadID: threadID,
		Body:     body,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, resource)
}
