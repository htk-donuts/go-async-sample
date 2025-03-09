package interactor

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/htk-donuts/go-async-sample/usecase/presenter"
	"github.com/htk-donuts/go-async-sample/usecase/repository"
)

type CSVInteractor interface {
	RequestCsvGenerate(ctx *gin.Context) error
}

type CSVInteractorImpl struct {
	repo      repository.ProductRepository
	presenter presenter.CSVPresenter
}

func NewCSVInteractor(repo repository.ProductRepository, presenter presenter.CSVPresenter) CSVInteractor {
	return &CSVInteractorImpl{repo: repo, presenter: presenter}
}

func (i *CSVInteractorImpl) RequestCsvGenerate(ctx *gin.Context) error {
	// コンテキストをコピーして非同期処理でも使用できるようにする
	// copyCtx := ctx.Copy()

	// 非同期でCSV生成処理を行う
	go func() {
		time.Sleep(5 * time.Second) // 重い処理をシミュレート
		products := i.repo.List(ctx)
		if err := i.presenter.OutputCSV(products); err != nil {
			log.Printf("Error outputting CSV: %v", err)
			return
		}
		log.Printf("CSV generation completed successfully")
	}()

	return nil
}
