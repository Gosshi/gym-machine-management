package route

import (
   "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    PingRoutes(r)
}
