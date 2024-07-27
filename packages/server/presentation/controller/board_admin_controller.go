package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BoardAdminController struct {
	boardApplicationService *service.BoardAdminApplicationService
}

func NewBoardAdminController(boardApplicationService *service.BoardAdminApplicationService) *BoardAdminController {
	return &BoardAdminController{boardApplicationService: boardApplicationService}
}

func (ctrl *BoardAdminController) FindAll(ctx *gin.Context) {
	var qs request.BoardAdminFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.boardApplicationService.FindAll(service.BoardAdminApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *BoardAdminController) Update(ctx *gin.Context) {
	var body request.BoardAdminUpdateRequest

	boardId, err := strconv.Atoi(ctx.Param("boardId"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.boardApplicationService.Update(service.BoardAdminApplicationServiceUpdateParams{
		Ctx:     context.Background(),
		BoardID: boardId,
		Body:    body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
