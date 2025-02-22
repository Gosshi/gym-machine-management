package route

import (
   "github.com/gin-gonic/gin"
)

// SetupRoutes ルーティングの設定
func SetupRoutes(r *gin.Engine) {
    PingRoutes(r)
}
