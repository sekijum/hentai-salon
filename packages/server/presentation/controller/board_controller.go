package controller

import (
	"context"
	"net/http"
	"server/application/service"
	request "server/presentation/request"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	boardService *service.BoardApplicationService
}

func NewBoardController(boardService *service.BoardApplicationService) *BoardController {
	return &BoardController{boardService: boardService}
}

func (ctrl *BoardController) FindAll(c *gin.Context) {
	boards, err := ctrl.boardService.FindAll(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "板の取得に失敗しました: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, boards)
}

func (ctrl *BoardController) Create(ginCtx *gin.Context) {
	var body request.BoardCreateRequest
	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.boardService.Create(context.Background(), ginCtx, body)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, nil)
}
