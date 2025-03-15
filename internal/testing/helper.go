package helper

import (
	"sync"
	"testing"
	"time"
)

// モック呼び出し回数のカウンター 主に非同期処理を伴うテストに利用
type MockCounter struct {
	mu             sync.Mutex
	counts         map[string]int
	expectedCounts map[string]int
}

func NewMockCounter(expectedCounts map[string]int) *MockCounter {
	return &MockCounter{
		counts:         make(map[string]int),
		expectedCounts: expectedCounts,
	}
}

// IncrementCount モック呼び出し回数をインクリメントする
func (c *MockCounter) IncrementCount(method string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counts[method]++
}

// AssertCounts モック呼び出し回数を期待呼び出し回数と比較してアサートする
func (c *MockCounter) AssertCounts(t *testing.T, maxRetry int, waitTime time.Duration) {
	AssertRetrying(t, func() bool {
		c.mu.Lock()
		allMatch := true

		for method, expectedCount := range c.expectedCounts {
			actualCount := c.counts[method]
			if actualCount != expectedCount {
				allMatch = false
				break
			}
		}
		c.mu.Unlock()

		return allMatch
	}, maxRetry, waitTime)
}

// AssertRetrying 指定回数リトライしてアサートを実行する
func AssertRetrying(t *testing.T, assertion func() bool,
	maxRetry int, waitTime time.Duration) {
	for i := 0; i < maxRetry; i++ {
		if assertion() {
			return
		}
		time.Sleep(waitTime)
	}
	t.Fail()
}
