package controller

import (
    "context"
    "net/http"
    "server/application/service"
    "server/infrastructure/datasource"
    "server/infrastructure/ent"

    "github.com/gin-gonic/gin"
)

type UserClientController struct {
    userAppService *service.UserApplicationClientService
}

func NewUserClientController(client *ent.Client) *UserClientController {
    userRepo := datasource.NewUserClientDatasource(client)
    userAppService := service.NewUserApplicationClientService(userRepo)
    return &UserClientController{userAppService: userAppService}
}

func (uc *UserClientController) GetUsers(ctx *gin.Context) {
    users, err := uc.userAppService.GetUsers(context.Background())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.JSON(http.StatusOK, users)
}
