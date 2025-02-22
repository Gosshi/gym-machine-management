package route

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gosshi/gym-machine-management/backend/controller"
)

func ManufacturerRoutes(app *fiber.App) {
    app.Post("/manufacturers", controller.CreateManufacturer)
}
