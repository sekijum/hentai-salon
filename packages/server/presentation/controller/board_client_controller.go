package board_client_controller

import (
	"context"
	"net/http"
	service "server/application/service"
	dto "server/presentation/dto/board"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

type BoardClientController struct {
	boardService *service.BoardClientService
}

func NewBoardClientController(boardService *service.BoardClientService) *BoardClientController {
	return &BoardClientController{boardService: boardService}
}

func (bc *BoardClientController) CreateBoard(ctx *gin.Context) {
	var req dto.BoardCreateClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	board, err := bc.boardService.CreateBoard(context.Background(), req.Title, req.Description, req.UserId, req.DefaultThreadTitle, req.DefaultThreadDescription)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, board)
}

var Provider = wire.NewSet(NewBoardClientController)
