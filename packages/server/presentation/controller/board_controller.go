package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	boardService *service.BoardApplicationService
}

func NewBoardController(boardService *service.BoardApplicationService) *BoardController {
	return &BoardController{boardService: boardService}
}

func (ctrl *BoardController) FindAll(ginCtx *gin.Context) {
	listResource, err := ctrl.boardService.FindAll(service.BoardApplicationServiceFindAllParams{
		Ctx: context.Background(),
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, listResource)
}

func (ctrl *BoardController) Create(ginCtx *gin.Context) {
	var body request.BoardCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resource, err := ctrl.boardService.Create(service.BoardApplicationServiceCreateParams{
		Ctx:    context.Background(),
		GinCtx: ginCtx,
		Body:   body,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, resource)
}
