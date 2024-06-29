package controller

import (
	"context"
	"net/http"
	"server/application/service"

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
