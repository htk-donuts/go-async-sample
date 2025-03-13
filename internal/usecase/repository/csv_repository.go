package repository

//go:generate mockgen -source=$GOFILE -destination=mock/$GOFILE

import (
	"context"

	"github.com/htk-donuts/go-async-sample/internal/domain/model"
)

type ProductRepository interface {
	List(ctx context.Context) []model.Product
}
