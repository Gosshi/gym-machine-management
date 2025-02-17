package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gosshi/gym-machine-management/docs" // ここで `docs` を import
	"github.com/gofiber/swagger"
)

func main() {
	// Fiber の初期化
	app := fiber.New()

	// Swagger のエンドポイントを追加
	app.Get("/swagger/*", swagger.HandlerDefault)

	// サーバーの起動
	port := ":8080"
	log.Println("🚀 Server is running on port", port)
	if err := app.Listen(port); err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
