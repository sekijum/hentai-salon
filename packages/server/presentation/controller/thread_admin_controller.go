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
	threadAdminApplicationService *service.ThreadAdminApplicationService
}

func NewThreadAdminController(threadAdminApplicationService *service.ThreadAdminApplicationService) *ThreadAdminController {
	return &ThreadAdminController{threadAdminApplicationService: threadAdminApplicationService}
}

func (ctrl *ThreadAdminController) FindAll(ctx *gin.Context) {
	var qs request.ThreadAdminFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.threadAdminApplicationService.FindAll(service.ThreadAdminApplicationServiceFindAllParams{
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
	threadID, err := strconv.Atoi(ctx.Param("threadID"))
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

	dto, err := ctrl.threadAdminApplicationService.FindByID(service.ThreadAdminApplicationServiceFindByIDParams{
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

	threadID, err := strconv.Atoi(ctx.Param("threadID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request.ThreadAdminUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.threadAdminApplicationService.Update(service.ThreadAdminApplicationServiceUpdateParams{
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
