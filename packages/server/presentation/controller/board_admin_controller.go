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

func (ctrl *BoardAdminController) FindAll(ginCtx *gin.Context) {
	var qs request.BoardAdminFindAllRequest

	if err := ginCtx.ShouldBindQuery(&qs); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ginCtx.GetInt("limit")
	qs.Offset = ginCtx.GetInt("offset")

	dto, err := ctrl.boardApplicationService.FindAll(service.BoardAdminApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, dto)
}

func (ctrl *BoardAdminController) Update(ginCtx *gin.Context) {
	var body request.BoardAdminUpdateRequest

	boardId, err := strconv.Atoi(ginCtx.Param("boardId"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid board ID"})
		return
	}

	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.boardApplicationService.Update(service.BoardAdminApplicationServiceUpdateParams{
		Ctx:     context.Background(),
		BoardID: boardId,
		Body:    body,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, dto)
}
