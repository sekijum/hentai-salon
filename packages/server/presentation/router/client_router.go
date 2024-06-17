package router

import (
    "server/infrastructure/ent"
    "server/presentation/controller"

    "github.com/gin-gonic/gin"
)

// SetupClientRoutes sets up the client routes
func SetupClientRoutes(r *gin.Engine, client *ent.Client) {
    userClientController := controller.NewUserClientController(client)
    clientRoutes := r.Group("/client")
    {
        clientRoutes.GET("/users", userClientController.GetUsers)
    }
}
