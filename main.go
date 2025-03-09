package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/htk-donuts/go-async-sample/infrastructure/repository"
	"github.com/htk-donuts/go-async-sample/interface/controller"
	"github.com/htk-donuts/go-async-sample/interface/presenter"
	"github.com/htk-donuts/go-async-sample/usecase/interactor"
)

func setupRouter() *gin.Engine {
	csvRepository := repository.NewProductRepository()
	csvPresenter := presenter.NewCSVPresenter()
	csvInteractor := interactor.NewCSVInteractor(csvRepository, csvPresenter)
	csvController := controller.NewCSVController(csvInteractor)

	router := gin.Default()

	router.POST("/generate-csv", csvController.HandleCSVGeneration)

	return router
}

func main() {
	router := setupRouter()
	port := ":8090"
	log.Printf("Server starting on port %s", port)
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}
