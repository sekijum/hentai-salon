package router

import (
	"server/di"
	"server/infrastructure/ent"

	"github.com/gin-gonic/gin"
)

func ClientRouter(r *gin.Engine, client *ent.Client) {
	boardController, err := di.InitializeClientControllers(client)
	if err != nil {
		panic(err)
	}

	clientRouter := r.Group("/client")
	{
		clientRouter.POST("/boards", boardController.CreateBoard)
	}
}
