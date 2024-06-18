package controller

import (
    "context"
    "net/http"
    "server/application/service"
    "server/infrastructure/datasource"
    "server/infrastructure/ent"

    "github.com/gin-gonic/gin"
)

type UserAdminController struct {
    userAppService *service.UserApplicationAdminService
}

func NewUserAdminController(client *ent.Client) *UserAdminController {
    userRepo := datasource.NewUserAdminDatasource(client)
    userAppService := service.NewUserApplicationAdminService(userRepo)
    return &UserAdminController{userAppService: userAppService}
}

func (uc *UserAdminController) GetUsers(ctx *gin.Context) {
    users, err := uc.userAppService.GetUsers(context.Background())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, users)
}
