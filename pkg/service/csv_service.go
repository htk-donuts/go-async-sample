package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/google/uuid"
)

// TaskStatus はCSV生成タスクのステータス情報です
type TaskStatus struct {
	Status   string // "processing", "completed", "failed"
	Progress int    // 0-100
	FilePath string
	FileName string
	Error    string
}

// CSVService はCSVファイルの生成を担当するサービスです
type CSVService struct {
	outputDir    string
	taskStatuses sync.Map // タスクID -> TaskStatus
}

// NewCSVService は新しいCSVServiceを作成します
func NewCSVService(outputDir string) *CSVService {
	// 出力ディレクトリが存在しない場合は作成
	if _, err := os.Stat(outputDir); os.IsNotExist(err) {
		os.MkdirAll(outputDir, 0755)
	}

	return &CSVService{
		outputDir:    outputDir,
		taskStatuses: sync.Map{},
	}
}

// StartCSVGeneration はCSV生成タスクを開始し、タスクIDを返します
func (s *CSVService) StartCSVGeneration(ctx context.Context, fileName string) string {
	taskID := uuid.New().String()

	// 初期ステータスを設定
	s.taskStatuses.Store(taskID, TaskStatus{
		Status:   "processing",
		Progress: 0,
		FileName: fileName,
	})

	// 非同期でCSV生成を開始
	go func() {
		filePath := filepath.Join(s.outputDir, fileName)
		err := s.generateCSV(ctx, taskID, filePath)

		if err != nil {
			s.taskStatuses.Store(taskID, TaskStatus{
				Status:   "failed",
				Progress: 0,
				Error:    err.Error(),
				FileName: fileName,
			})
		} else {
			s.taskStatuses.Store(taskID, TaskStatus{
				Status:   "completed",
				Progress: 100,
				FilePath: filePath,
				FileName: fileName,
			})
		}
	}()

	return taskID
}

// GetCSVGenerationStatus はタスクIDに対応するCSV生成状況を返します
func (s *CSVService) GetCSVGenerationStatus(taskID string) TaskStatus {
	value, exists := s.taskStatuses.Load(taskID)
	if !exists {
		return TaskStatus{
			Status: "not_found",
			Error:  "指定されたタスクIDが見つかりません",
		}
	}

	return value.(TaskStatus)
}

// generateCSV はCSVファイルを生成します
func (s *CSVService) generateCSV(ctx context.Context, taskID, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("ファイル作成エラー: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// ヘッダーを書き込み
	header := []string{"ID", "Name", "Email", "Date", "Value"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("ヘッダー書き込みエラー: %v", err)
	}

	// サンプルデータ行数
	totalRows := 1000

	// データを生成して書き込み
	for i := 1; i <= totalRows; i++ {
		// コンテキストのキャンセル確認
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			// 処理続行
		}

		// CSV行を生成
		row := []string{
			fmt.Sprintf("%d", i),
			fmt.Sprintf("User %d", i),
			fmt.Sprintf("user%d@example.com", i),
			time.Now().Format(time.RFC3339),
			fmt.Sprintf("%.2f", float64(i)*1.5),
		}

		if err := writer.Write(row); err != nil {
			return fmt.Errorf("行書き込みエラー: %v", err)
		}

		// 進捗状況を更新 (10行ごと)
		if i%10 == 0 {
			progress := (i * 100) / totalRows
			s.taskStatuses.Store(taskID, TaskStatus{
				Status:   "processing",
				Progress: progress,
				FileName: filepath.Base(filePath),
			})
		}

		// 生成に時間がかかっていることをシミュレート
		time.Sleep(5 * time.Millisecond)
	}

	return nil
}
