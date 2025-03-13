package presenter

import (
	"encoding/csv"
	"os"

	"github.com/htk-donuts/go-async-sample/internal/domain/model"
)

type CSVPresenterImpl struct{}

func NewCSVPresenter() *CSVPresenterImpl {
	return &CSVPresenterImpl{}
}

func (p *CSVPresenterImpl) OutputCSV(products []model.Product) error {
	file, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	headers := []string{"商品名", "価格", "在庫数"}
	if err := writer.Write(headers); err != nil {
		return err
	}

	for _, product := range products {
		row := []string{product.Name, product.Price, product.Stock}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
