# Go Async CSV Generator with Gin

## 概要
このプロジェクトは、Go言語とGinフレームワークを使用して、非同期でCSVファイルを生成するAPIを実装したサンプルです。
コンテキストを活用してリクエストのキャンセルやタイムアウトに対応しています。

## 主な機能
- 非同期CSV生成
- タスク進捗状況の確認
- 生成されたCSVファイルのダウンロード
- コンテキストを使用したキャンセル対応
- Ginフレームワークによるルーティング

## APIエンドポイント

### ヘルスチェック
```
GET /health
```

### CSV生成開始
```
POST /api/csv/generate?filename=report.csv
```
レスポンス:
```json
{
  "status": "processing",
  "task_id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
  "message": "CSV生成を開始しました。ステータス確認エンドポイントで進捗を確認できます。"
}
```

### タスク状態確認
```
GET /api/csv/status/{taskId}
```
レスポンス例 (処理中):
```json
{
  "status": "processing",
  "progress": 45,
  "message": "CSV生成中です"
}
```

レスポンス例 (完了):
```json
{
  "status": "completed",
  "file_path": "/path/to/file.csv",
  "message": "CSV生成が完了しました"
}
```

### ファイルダウンロード
```
GET /api/csv/download/{taskId}
```
成功した場合、CSVファイルがダウンロードされます。

## 使用方法

### サーバーの起動
```bash
go run cmd/main.go
```

### CSVファイル生成のリクエスト例
```bash
curl -X POST "http://localhost:8080/api/csv/generate?filename=my_report.csv"
```

### タスク状態確認の例
```bash
curl http://localhost:8080/api/csv/status/{taskId}
```

### ファイルダウンロードの例
```bash
curl -OJ http://localhost:8080/api/csv/download/{taskId}
```
