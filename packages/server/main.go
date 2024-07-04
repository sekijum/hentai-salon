package main

import (
	"time"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	presentation "server/presentation"
	"server/shared/di"
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

	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://192.168.10.30:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}

	r.Use(cors.New(config))

	r.SetTrustedProxies([]string{"127.0.0.1"})

	presentation.SetupRouter(r, controllers)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
