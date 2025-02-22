package route

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	// Gin のテストモードに設定
	gin.SetMode(gin.TestMode)

	// テスト用のエンジンを作成
	router := gin.Default()
	// /ping エンドポイントを登録（SetupRoutesが登録している前提）
	router.GET("/ping", Ping)

	// テストリクエストを作成
	req, _ := http.NewRequest("GET", "/ping", nil)
	w := httptest.NewRecorder()

	// ルーターを通してリクエストを実行
	router.ServeHTTP(w, req)

	// ステータスコードとレスポンスの内容を検証
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "pong")
}
