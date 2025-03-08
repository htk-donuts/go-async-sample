package repository

import "github.com/htk-donuts/go-async-sample/domain/model"

type CSVRepositoryImpl struct{}

func NewCSVRepository() *CSVRepositoryImpl {
	return &CSVRepositoryImpl{}
}

func (r *CSVRepositoryImpl) GetProducts() []model.Product {
	return []model.Product{
		{Name: "商品A", Price: "1000", Stock: "50"},
		{Name: "商品B", Price: "2000", Stock: "30"},
		{Name: "商品C", Price: "1500", Stock: "100"},
		{Name: "商品D", Price: "3000", Stock: "20"},
	}
}
