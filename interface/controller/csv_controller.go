package controller

import (
	"log"
	"net/http"

	"github.com/htk-donuts/go-async-sample/usecase/interactor"
)

type CSVController struct {
	interactor *interactor.CSVInteractor
}

func NewCSVController(interactor *interactor.CSVInteractor) *CSVController {
	return &CSVController{interactor: interactor}
}

func (c *CSVController) HandleCSVGeneration(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	go func() {
		if err := c.interactor.GenerateCSV(); err != nil {
			log.Printf("CSV generation error: %v", err)
			return
		}
	}()

	w.WriteHeader(http.StatusAccepted)
}
