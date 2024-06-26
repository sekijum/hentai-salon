package main

import (
	"log"
	"os"

	presentation "server/presentation"
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers, cleanup, err := di.InitializeControllers()
	if err != nil {
		log.Fatalf("failed to initialize controllers: %v", err)
	}
	defer cleanup()

	if os.Getenv("GIN_MODE") != "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.SetTrustedProxies([]string{"127.0.0.1"})

	presentation.SetupRouter(r, controllers)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
