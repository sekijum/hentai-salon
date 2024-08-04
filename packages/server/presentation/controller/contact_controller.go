package controller

import (
	"net/http"
	"server/application/service"
	"server/presentation/request"

	"github.com/gin-gonic/gin"
)

type ContactController struct {
	contactApplicationService *service.ContactApplicationService
}

func NewContactController(contactApplicationService *service.ContactApplicationService) *ContactController {
	return &ContactController{contactApplicationService: contactApplicationService}
}

func (ctrl *ContactController) Create(ctx *gin.Context) {
	var body request.ContactCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctrl.contactApplicationService.Create(service.ContactApplicationServiceCreateParams{
		Ctx:      ctx.Request.Context(),
		ClientIP: ctx.ClientIP(),
		Body:     body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
