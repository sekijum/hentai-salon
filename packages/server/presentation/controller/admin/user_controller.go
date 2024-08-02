package controller_admin

import (
	"context"
	"net/http"
	service_admin "server/application/service/admin"
	request_admin "server/presentation/request/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userApplicationService *service_admin.UserApplicationService
}

func NewUserController(userApplicationService *service_admin.UserApplicationService) *UserController {
	return &UserController{userApplicationService: userApplicationService}
}

func (ctrl *UserController) FindByID(ctx *gin.Context) {
	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.userApplicationService.FindByID(service_admin.UserApplicationServiceFindByIDParams{
		Ctx:    context.Background(),
		UserID: userID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserController) FindAll(ctx *gin.Context) {
	var qs request_admin.UserFindAllRequest

	if err := ctx.ShouldBindQuery(&qs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ctx.GetInt("limit")
	qs.Offset = ctx.GetInt("offset")

	dto, err := ctrl.userApplicationService.FindAll(service_admin.UserApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, dto)
}

func (ctrl *UserController) Update(ctx *gin.Context) {

	userID, err := strconv.Atoi(ctx.Param("userID"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var body request_admin.UserUpdateRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dto, err := ctrl.userApplicationService.Update(service_admin.UserApplicationServiceUpdateParams{
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
