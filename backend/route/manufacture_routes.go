package route

import (
    "github.com/gin-gonic/gin"
    "github.com/gosshi/gym-machine-management/backend/controller"
)

func ManufacturerRoutes(r *gin.Engine) {
    r.POST("/manufacturers", controller.CreateManufacturer)
}
