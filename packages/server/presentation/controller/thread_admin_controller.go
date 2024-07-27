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

func (ctrl *ThreadAdminController) FindAll(ctx *gin.Context) {
	var qs request.ThreadAdminFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.threadApplicationService.FindAll(service.ThreadAdminApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadAdminController) FindByID(ctx *gin.Context) {
	threadID, err := strconv.Atoi(ctx.Param("threadId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var qs request.ThreadAdminFindByIDRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.threadApplicationService.FindByID(service.ThreadAdminApplicationServiceFindByIDParams{
		Ctx:      context.Background(),
		ThreadID: threadID,
		Qs:       qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ThreadAdminController) Update(ctx *gin.Context) {
	var body request.ThreadAdminUpdateRequest

	threadID, err := strconv.Atoi(ctx.Param("threadId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.threadApplicationService.Update(service.ThreadAdminApplicationServiceUpdateParams{
		Ctx:      context.Background(),
		ThreadID: threadID,
		Body:     body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
