package controller

import (
	"context"
	"net/http"
	"server/application/service"

	"github.com/gin-gonic/gin"
)

type AdController struct {
	AdService *service.AdApplicationService
}

func NewAdController(AdService *service.AdApplicationService) *AdController {
	return &AdController{AdService: AdService}
}

func (ctrl *AdController) FindAll(ctx *gin.Context) {
	dto, err := ctrl.AdService.FindAll(service.AdApplicationServiceFindAllParams{
		Ctx: context.Background(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
