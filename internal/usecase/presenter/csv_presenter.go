package presenter

//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE

import "github.com/htk-donuts/go-async-sample/internal/domain/model"

type CSVPresenter interface {
	OutputCSV([]model.Product) error
}
