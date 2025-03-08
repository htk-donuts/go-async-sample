package repository

import "github.com/htk-donuts/go-async-sample/domain/model"

type CSVRepository interface {
	GetProducts() []model.Product
}
