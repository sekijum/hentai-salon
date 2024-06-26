package board_client_controller

import (
	"context"
	"net/http"
	service "server/application/service"
	request "server/presentation/request/board"

	"github.com/gin-gonic/gin"
)

type BoardClientController struct {
	boardClientService *service.BoardClientService
}

func NewBoardClientController(boardClientService *service.BoardClientService) *BoardClientController {
	return &BoardClientController{boardClientService: boardClientService}
}

func (ctrl *BoardClientController) Create(ginCtx *gin.Context) {
	var body request.BoardCreateClientRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := ctrl.boardClientService.Create(context.Background(), ginCtx, body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, nil)
}
