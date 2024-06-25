package main

import (
	"log"

	routes "server/presentation"
	"server/shared/di"

	"github.com/gin-gonic/gin"
)

func main() {
	controllers, cleanup, err := di.InitializeControllers()
	if err != nil {
		log.Fatalf("failed to initialize controllers: %v", err)
	}
	defer cleanup()

	// Ginルーターを新しく作成
	r := gin.Default()

	// ルートを設定
	routes.SetupRouter(r, controllers)

	// サーバーを開始
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
