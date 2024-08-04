package controller_admin

import (
	"context"
	"net/http"
	service_admin "server/application/service/admin"
	request_admin "server/presentation/request/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	contactApplicationService *service_admin.ContactApplicationService
}

func NewContactController(contactApplicationService *service_admin.ContactApplicationService) *ContactController {
	return &ContactController{contactApplicationService: contactApplicationService}
}

func (ctrl *ContactController) FindByID(ctx *gin.Context) {
	contactID, err := strconv.Atoi(ctx.Param("contactID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.contactApplicationService.FindByID(service_admin.ContactApplicationServiceFindByIDParams{
		Ctx:       context.Background(),
		ContactID: contactID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ContactController) FindAll(ctx *gin.Context) {
	var qs request_admin.ContactFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.contactApplicationService.FindAll(service_admin.ContactApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *ContactController) UpdateStatus(ctx *gin.Context) {
	var body request_admin.ContactUpdateStatusRequest

	ContactID, err := strconv.Atoi(ctx.Param("contactID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.contactApplicationService.UpdateStatus(service_admin.ContactApplicationServiceUpdateParams{
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
