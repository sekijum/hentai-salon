package controller

import (
	"context"
	"net/http"
	applicationService "server/application/service"
	request "server/presentation/request"

	"github.com/gin-gonic/gin"
)

type BoardAdminController struct {
	boardAdminApplicationService *applicationService.BoardAdminApplicationService
}

func NewBoardAdminController(boardAdminApplicationService *applicationService.BoardAdminApplicationService) *BoardAdminController {
	return &BoardAdminController{boardAdminApplicationService: boardAdminApplicationService}
}

func (ctrl *BoardAdminController) Create(ginCtx *gin.Context) {
	var body request.BoardCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.boardAdminApplicationService.Create(context.Background(), ginCtx, body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, nil)
}
