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

func (ctrl *UserAdminController) FindAll(ginCtx *gin.Context) {
	var qs request.UserAdminFindAllRequest

	if err := ginCtx.ShouldBindQuery(&qs); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qs.Limit = ginCtx.GetInt("limit")
	qs.Offset = ginCtx.GetInt("offset")

	listResource, err := ctrl.userAdminApplicationService.FindAll(service.UserAdminApplicationServiceFindAllParams{
		Ctx: context.Background(),
		Qs:  qs,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, listResource)
}

func (ctrl *UserAdminController) Update(ginCtx *gin.Context) {
	var body request.UserAdminUpdateRequest

	userID, err := strconv.Atoi(ginCtx.Param("userId"))
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := ginCtx.ShouldBindJSON(&body); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResource, err := ctrl.userAdminApplicationService.Update(service.UserAdminApplicationServiceUpdateParams{
		Ctx:    context.Background(),
		UserID: userID,
		Body:   body,
	})
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, userResource)
}
