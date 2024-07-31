package controller

import (
	"context"
	"net/http"
	"server/application/service"
	"server/presentation/request"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAdminController struct {
	userAdminApplicationService *service.UserAdminApplicationService
}

func NewUserAdminController(userAdminApplicationService *service.UserAdminApplicationService) *UserAdminController {
	return &UserAdminController{userAdminApplicationService: userAdminApplicationService}
}

func (ctrl *UserAdminController) FindByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.userAdminApplicationService.FindByID(service.UserAdminApplicationServiceFindByIDParams{
		Ctx:    context.Background(),
		UserID: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserAdminController) FindAll(ctx *gin.Context) {
	var qs request.UserAdminFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.userAdminApplicationService.FindAll(service.UserAdminApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserAdminController) Update(ctx *gin.Context) {
	var body request.UserAdminUpdateRequest

	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.userAdminApplicationService.Update(service.UserAdminApplicationServiceUpdateParams{
		Ctx:    context.Background(),
		UserID: userID,
		Body:   body,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}
