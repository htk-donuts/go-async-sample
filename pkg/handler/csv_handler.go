package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hikarutakao/go-async-sample/pkg/service"
)

// CSVHandler は CSV生成処理のためのハンドラです
type CSVHandler struct {
	csvService *service.CSVService
}

// NewCSVHandler は新しいCSVHandlerを作成します
func NewCSVHandler(csvService *service.CSVService) *CSVHandler {
	return &CSVHandler{
		csvService: csvService,
	}
}

// GenerateCSV はCSVファイルを非同期で生成するエンドポイントです
func (h *CSVHandler) GenerateCSV(c *gin.Context) {
	// リクエストパラメータを取得
	fileName := c.Query("filename")
	if fileName == "" {
		fileName = "report.csv" // デフォルトファイル名
	}

	// タスクの開始レスポンスを返す
	taskID := h.csvService.StartCSVGeneration(c.Request.Context(), fileName)

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "processing",
		"task_id": taskID,
		"message": "CSV生成を開始しました。ステータス確認エンドポイントで進捗を確認できます。",
	})
}

// GetCSVStatus はCSV生成タスクの状態を確認するエンドポイントです
func (h *CSVHandler) GetCSVStatus(c *gin.Context) {
	taskID := c.Param("taskId")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "タスクIDが必要です",
		})
		return
	}

	status := h.csvService.GetCSVGenerationStatus(taskID)

	switch status.Status {
	case "completed":
		c.JSON(http.StatusOK, gin.H{
			"status":    status.Status,
			"file_path": status.FilePath,
			"message":   "CSV生成が完了しました",
		})
	case "failed":
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  status.Status,
			"error":   status.Error,
			"message": "CSV生成に失敗しました",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"status":   status.Status,
			"progress": status.Progress,
			"message":  "CSV生成中です",
		})
	}
}

// DownloadCSV は生成されたCSVファイルをダウンロードするエンドポイントです
func (h *CSVHandler) DownloadCSV(c *gin.Context) {
	taskID := c.Param("taskId")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "タスクIDが必要です",
		})
		return
	}

	status := h.csvService.GetCSVGenerationStatus(taskID)

	if status.Status != "completed" {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "指定されたタスクのCSVファイルがまだ生成されていません",
		})
		return
	}

	// ファイル転送
	c.FileAttachment(status.FilePath, status.FileName)
}
