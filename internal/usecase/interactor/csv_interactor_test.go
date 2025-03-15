package interactor

import (
	"context"
	"errors"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/htk-donuts/go-async-sample/internal/domain/model"
	helper "github.com/htk-donuts/go-async-sample/internal/testing"
	mockPresenter "github.com/htk-donuts/go-async-sample/internal/usecase/presenter/mock"
	mockRepository "github.com/htk-donuts/go-async-sample/internal/usecase/repository/mock"
	"go.uber.org/mock/gomock"
)

// 一定時間待機させて非同期処理完了を待つテストケース。
func TestRequestCsvGenerate(t *testing.T) {
	// テストケースの定義
	tests := []struct {
		name          string
		setupMocks    func(*mockRepository.MockProductRepository, *mockPresenter.MockCSVPresenter)
		expectedError error
	}{
		{
			name: "正常系: CSV生成が成功する",
			setupMocks: func(repo *mockRepository.MockProductRepository, presenter *mockPresenter.MockCSVPresenter) {
				products := []model.Product{
					{Name: "商品A", Price: "1000", Stock: "50"},
				}
				repo.EXPECT().List(gomock.Any()).Return(products)
				presenter.EXPECT().OutputCSV(products).Return(nil)
			},
			expectedError: nil,
		},
		{
			name: "異常系: CSV出力でエラーが発生",
			setupMocks: func(repo *mockRepository.MockProductRepository, presenter *mockPresenter.MockCSVPresenter) {
				products := []model.Product{
					{Name: "商品A", Price: "1000", Stock: "50"},
				}
				repo.EXPECT().List(gomock.Any()).Return(products)
				presenter.EXPECT().OutputCSV(products).Return(errors.New("CSV出力エラー"))
			},
			expectedError: nil, // 非同期処理なのでエラーは返らない
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックコントローラーの作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// モックの作成
			mockProductRepo := mockRepository.NewMockProductRepository(ctrl)
			mockCsvPresenter := mockPresenter.NewMockCSVPresenter(ctrl)

			// モックの設定
			tt.setupMocks(mockProductRepo, mockCsvPresenter)

			// テスト対象のインスタンス作成
			interactor := NewCSVInteractor(mockProductRepo, mockCsvPresenter)

			// テスト用のgin.Context作成
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// テスト実行
			err := interactor.RequestCsvGenerate(c)

			// 非同期処理の完了を待つ 指定時間以内に終わる担保はどこにもなくFlakyになる可能性がある。過剰待機になる可能性もありCI実行時間の無駄も発生。
			time.Sleep(6 * time.Second)

			// アサーション
			if err != tt.expectedError {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}

// モック関数の呼び出し回数をカウントするテストケース。
func TestRequestCsvGenerateCount(t *testing.T) {
	const ( // モック関数の呼び出し回数をカウントするためのKeyとして利用する文字列
		mockMethodProductRepositoryList string = "ProductRepositoryList"
		mockMethodCSVPresenterOutputCSV string = "CSVPresenterOutputCSV"
	)

	tests := []struct {
		name          string
		setupMocks    func(*mockRepository.MockProductRepository, *mockPresenter.MockCSVPresenter) *helper.MockCounter
		expectedError error
	}{
		{
			name: "正常系: CSV生成が成功する",
			setupMocks: func(repo *mockRepository.MockProductRepository, presenter *mockPresenter.MockCSVPresenter) *helper.MockCounter {
				mockCounter := helper.NewMockCounter(
					map[string]int{
						mockMethodProductRepositoryList: 1,
						mockMethodCSVPresenterOutputCSV: 1,
					})
				products := []model.Product{
					{Name: "商品A", Price: "1000", Stock: "50"},
				}
				repo.EXPECT().List(gomock.Any()).
					DoAndReturn(func(ctx context.Context) ([]model.Product, error) {
						mockCounter.IncrementCount(mockMethodProductRepositoryList)
						return products, nil
					})
				presenter.EXPECT().OutputCSV(products).Return(nil).DoAndReturn(
					func(products []model.Product) error {
						mockCounter.IncrementCount(mockMethodCSVPresenterOutputCSV)
						return nil
					})
				return mockCounter
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックコントローラーの作成
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// モックの作成
			mockProductRepo := mockRepository.NewMockProductRepository(ctrl)
			mockCsvPresenter := mockPresenter.NewMockCSVPresenter(ctrl)

			// モックの設定
			mockCounter := tt.setupMocks(mockProductRepo, mockCsvPresenter)

			// テスト対象のインスタンス作成
			interactor := NewCSVInteractor(mockProductRepo, mockCsvPresenter)

			// テスト用のgin.Context作成
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// テスト実行
			err := interactor.RequestCsvGenerate(c)

			// アサーション
			if err != tt.expectedError {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}

			// 非同期処理アサーション モッキングした関数が想定回数呼ばれることを確認
			// 1秒ごとにチェックを行い、10回以内に想定回数呼ばれることを確認
			mockCounter.AssertCounts(t, 10, 1*time.Second)
		})
	}
}
