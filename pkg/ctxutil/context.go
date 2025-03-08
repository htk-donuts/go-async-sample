package ctxutil

import (
	"context"
	"time"
)

// CreateTimeoutContext はタイムアウト付きのコンテキストを作成します
func CreateTimeoutContext(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration)
}

// CreateCancelableContext はキャンセル可能なコンテキストを作成します
func CreateCancelableContext() (context.Context, context.CancelFunc) {
	return context.WithCancel(context.Background())
}

// CreateChildContext は親コンテキストから子コンテキストを作成します
func CreateChildContext(parent context.Context) (context.Context, context.CancelFunc) {
	return context.WithCancel(parent)
}
