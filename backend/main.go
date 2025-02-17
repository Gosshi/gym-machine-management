package main

import (
	"log"
	"github.com/gofiber/fiber/v2"

    "github.com/gin-gonic/gin"
	_ "github.com/gosshi/gym-machine-management/docs" // Swagger のドキュメントをインポート
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Fiber の初期化
	app := fiber.New()


	// サーバーの起動
	port := ":8080"
	log.Println("🚀 Server is running on port", port)
	if err := app.Listen(port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
