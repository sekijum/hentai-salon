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

func (ctrl *BoardController) FindAll(ctx *gin.Context) {
	dto, err := ctrl.boardService.FindAll(service.BoardApplicationServiceFindAllParams{
		Ctx: context.Background(),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *BoardController) Create(ctx *gin.Context) {
	var body request.BoardCreateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := ctx.Get("userID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "ユーザーIDがコンテキストに存在しません"})
		return
	}

	dto, err := ctrl.boardService.Create(service.BoardApplicationServiceCreateParams{
		Ctx:    context.Background(),
		UserID: userID.(int),
		Body:   body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
