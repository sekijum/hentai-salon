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
	contactAdminApplicationService *service.ContactAdminApplicationService
}

func NewContactAdminController(contactAdminApplicationService *service.ContactAdminApplicationService) *ContactAdminController {
	return &ContactAdminController{contactAdminApplicationService: contactAdminApplicationService}
}

func (ctrl *ContactAdminController) FindByID(ctx *gin.Context) {
	contactID, err := strconv.Atoi(ctx.Param("contactID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.contactAdminApplicationService.FindByID(service.ContactAdminApplicationServiceFindByIDParams{
		Ctx:       context.Background(),
		ContactID: contactID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ContactAdminController) FindAll(ctx *gin.Context) {
	var qs request.ContactAdminFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.contactAdminApplicationService.FindAll(service.ContactAdminApplicationServiceFindAllParams{
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

	ContactID, err := strconv.Atoi(ctx.Param("contactID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.contactAdminApplicationService.Update(service.ContactAdminApplicationServiceUpdateParams{
		Ctx:       context.Background(),
		ContactID: ContactID,
		Body:      body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
