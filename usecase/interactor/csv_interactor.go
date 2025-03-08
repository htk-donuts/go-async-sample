package interactor

//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE

import (
	"log"
	"time"

	"github.com/htk-donuts/go-async-sample/usecase/presenter"
	"github.com/htk-donuts/go-async-sample/usecase/repository"
)

type CSVInteractor struct {
	repo      repository.CSVRepository
	presenter presenter.CSVPresenter
}

func NewCSVInteractor(repo repository.CSVRepository, presenter presenter.CSVPresenter) *CSVInteractor {
	return &CSVInteractor{repo: repo, presenter: presenter}
}

func (i *CSVInteractor) GenerateCSV() error {
	products := i.repo.GetProducts()
	time.Sleep(5 * time.Second) // 重い処理をシミュレート
	if err := i.presenter.OutputCSV(products); err != nil {
		return err
	}
	log.Printf("CSV generation completed successfully")
	return nil
}
