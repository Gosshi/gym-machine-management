package route

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gosshi/gym-machine-management/backend/route/pingRoutes"

)

func SetupRoutes(app *fiber.App) {
    pingRoutes.PingRoutes(app)

}
