package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
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
