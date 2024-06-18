package router

import (
    "server/infrastructure/ent"
    "server/presentation/controller"

    "github.com/gin-gonic/gin"
)

// SetupAdminRoutes sets up the admin routes
func SetupAdminRoutes(r *gin.Engine, client *ent.Client) {
    userAdminController := controller.NewUserAdminController(client)
    admin := r.Group("/admin")
    {
        admin.GET("/users", userAdminController.GetUsers)
    }
}
