package main

import (
	"log"
	"net/http"

	"github.com/htk-donuts/go-async-sample/infrastructure/repository"
	"github.com/htk-donuts/go-async-sample/interface/controller"
	"github.com/htk-donuts/go-async-sample/interface/presenter"
	"github.com/htk-donuts/go-async-sample/usecase/interactor"
)

func setupServer() *http.ServeMux {
	csvRepository := repository.NewCSVRepository()
	csvPresenter := presenter.NewCSVPresenter()
	csvInteractor := interactor.NewCSVInteractor(csvRepository, csvPresenter)
	csvController := controller.NewCSVController(csvInteractor)

	mux := http.NewServeMux()
	mux.HandleFunc("/generate-csv", csvController.HandleCSVGeneration)

	return mux
}

func main() {
	mux := setupServer()
	port := ":8090"
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
