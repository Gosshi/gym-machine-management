package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/gosshi/gym-machine-management/docs" // swag によって生成される docs パッケージ
	"github.com/gosshi/gym-machine-management/route"

	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

func main() {
	r := gin.Default()

	// ルーティング
	route.SetupRoutes(r)

	// Swagger UI を /swagger/*any で提供
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// サーバーの起動
	port := ":8080"

	log.Println("🚀 Server is running on port", port)
	if err := r.Run(port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}