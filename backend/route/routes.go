package route

import (
    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
    ManufacturerRoutes(app) // メーカー関連のルートを登録
}
