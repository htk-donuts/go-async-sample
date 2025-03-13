package repository

import (
	"context"

	"github.com/htk-donuts/go-async-sample/internal/domain/model"
)

type ProductRepositoryImpl struct{}

func NewProductRepository() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{}
}

func (r *ProductRepositoryImpl) List(ctx context.Context) []model.Product {
	return []model.Product{
		{Name: "商品A", Price: "1000", Stock: "50"},
		{Name: "商品B", Price: "2000", Stock: "30"},
		{Name: "商品C", Price: "1500", Stock: "100"},
		{Name: "商品D", Price: "3000", Stock: "20"},
	}
}
