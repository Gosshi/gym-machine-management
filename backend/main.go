package main

import (
    "github.com/gofiber/fiber/v2"
)

func main() {
    // Fiber のインスタンスを作成
    app := fiber.New()

    // ルートエンドポイントの定義
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, Fiber!")
    })

    // サーバーをポート3000で起動
    app.Listen(":3000")
}
