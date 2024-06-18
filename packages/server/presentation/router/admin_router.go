package router

import (
    "server/infrastructure/ent"
    "server/presentation/controller"

    "github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.Engine, client *ent.Client) {
    userAdminController := controller.NewUserAdminController(client)
    admin := r.Group("/admin")
    {
        admin.GET("/users", userAdminController.GetUsers)
    }
}
