package route

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

// PingHandler シンプルなヘルスチェック API
// @Summary Pingチェック
// @Description サーバーが動作しているか確認
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string "サーバーの状態を返す"
// @Router /ping [get]
// PingRoutes Ping API のルート設定
func PingRoutes(r *gin.Engine) {
    r.GET("/ping", Ping)
}

func Ping(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

