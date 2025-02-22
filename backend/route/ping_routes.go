package route

import (
	"github.com/gin-gonic/gin"
	"github.com/gosshi/gym-machine-management/controller"
)

// PingHandler シンプルなヘルスチェック API
// @Summary Pingチェック
// @Description サーバーが動作しているか確認
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string "サーバーの状態を返す"
// @Router /ping [get]
// PingRoutes Ping API のルート設定
func PingRoutes(app *fiber.App) {
    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "pong"})
    })
}

