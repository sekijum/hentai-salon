package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactAdminController struct {
	ContactAdminApplicationService *service.ContactAdminApplicationService
}

func NewContactAdminController(ContactAdminApplicationService *service.ContactAdminApplicationService) *ContactAdminController {
	return &ContactAdminController{ContactAdminApplicationService: ContactAdminApplicationService}
}

func (ctrl *ContactAdminController) FindAll(ctx *gin.Context) {
	var qs request.ContactAdminFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.ContactAdminApplicationService.FindAll(service.ContactAdminApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ContactAdminController) Update(ctx *gin.Context) {
	var body request.ContactAdminUpdateRequest

	ContactId, err := strconv.Atoi(ctx.Param("ContactId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.ContactAdminApplicationService.Update(service.ContactAdminApplicationServiceUpdateParams{
		Ctx:       context.Background(),
		ContactID: ContactId,
		Body:      body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
