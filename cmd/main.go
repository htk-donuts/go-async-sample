package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hikarutakao/go-async-sample/pkg/handler"
	"github.com/hikarutakao/go-async-sample/pkg/middleware"
	"github.com/hikarutakao/go-async-sample/pkg/service"
)

func main() {
	// 作業ディレクトリを取得
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("作業ディレクトリの取得に失敗: %v", err)
	}

	// CSV出力ディレクトリを作成
	outputDir := filepath.Join(workDir, "output")
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, 0755)
	}

	// サービスを初期化
	csvService := service.NewCSVService(outputDir)

	// ハンドラを初期化
	csvHandler := handler.NewCSVHandler(csvService)

	// Ginルーターを設定
	router := setupRouter(csvHandler)

	// サーバー作成
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// サーバーを非同期で起動
	go func() {
		fmt.Println("サーバーを起動しています (http://localhost:8080)...")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("サーバー起動エラー: %v", err)
		}
	}()

	// シグナル処理のためのチャネル
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("サーバーをシャットダウンしています...")

	// グレースフルシャットダウンのためのコンテキスト
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("サーバーの強制シャットダウン: %v", err)
	}

	fmt.Println("サーバーを正常にシャットダウンしました")
}

// setupRouter はGinルーターを設定して返します
func setupRouter(csvHandler *handler.CSVHandler) *gin.Engine {
	router := gin.Default()

	// ミドルウェアを適用
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.TimeoutMiddleware(60 * time.Second))

	// ヘルスチェックエンドポイント
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
			"time":   time.Now().Format(time.RFC3339),
		})
	})

	// CSV生成 API
	api := router.Group("/api")
	{
		api.POST("/csv/generate", csvHandler.GenerateCSV)
		api.GET("/csv/status/:taskId", csvHandler.GetCSVStatus)
		api.GET("/csv/download/:taskId", csvHandler.DownloadCSV)
	}

	return router
}
