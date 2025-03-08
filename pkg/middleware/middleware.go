package middleware

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// TimeoutMiddleware はリクエストのタイムアウトを設定するミドルウェアです
func TimeoutMiddleware(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		// タイムアウト付きのコンテキストを作成
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()

		// リクエストのコンテキストを更新
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// LoggerMiddleware はリクエストのログを記録するミドルウェアです
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエスト開始時間
		startTime := time.Now()

		// リクエストパスを取得
		path := c.Request.URL.Path

		// ハンドラーを実行
		c.Next()

		// リクエスト終了時間と処理時間
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// ステータスコード
		statusCode := c.Writer.Status()

		// ログ出力
		fmt.Printf("[%s] %s %s %d %v\n",
			endTime.Format("2006/01/02 15:04:05"),
			c.Request.Method,
			path,
			statusCode,
			latency,
		)
	}
}
