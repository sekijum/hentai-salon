package main

import (
	"log"
	"os"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"

	presentation "server/presentation"
	"server/presentation/middleware"
	"server/shared/di"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	controllers, cleanup, err := di.InitializeControllers()
	if err != nil {
		log.Fatalf("failed to initialize controllers: %v", err)
	}
	defer cleanup()

	if os.Getenv("GIN_MODE") != "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.Use(middleware.CORS())

	presentation.SetupRouter(r, controllers)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
