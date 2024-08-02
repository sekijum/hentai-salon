package controller_admin

import (
	"context"
	"net/http"
	service_admin "server/application/service/admin"
	request_admin "server/presentation/request/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardController struct {
	boardApplicationService *service_admin.BoardApplicationService
}

func NewBoardController(boardApplicationService *service_admin.BoardApplicationService) *BoardController {
	return &BoardController{boardApplicationService: boardApplicationService}
}

func (ctrl *BoardController) FindByID(ctx *gin.Context) {
	boardID, err := strconv.Atoi(ctx.Param("boardID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.boardApplicationService.FindByID(service_admin.BoardApplicationServiceFindByIDParams{
		Ctx:     context.Background(),
		BoardID: boardID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *BoardController) FindAll(ctx *gin.Context) {
	var qs request_admin.BoardFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.boardApplicationService.FindAll(service_admin.BoardApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *BoardController) Update(ctx *gin.Context) {
	var body request_admin.BoardUpdateRequest

	boardID, err := strconv.Atoi(ctx.Param("boardID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.boardApplicationService.Update(service_admin.BoardApplicationServiceUpdateParams{
		Ctx:     context.Background(),
		BoardID: boardID,
		Body:    body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
