package presenter

import "github.com/htk-donuts/go-async-sample/internal/domain/model"

type CSVPresenter interface {
	OutputCSV([]model.Product) error
}
