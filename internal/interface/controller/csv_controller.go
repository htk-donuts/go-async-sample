package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htk-donuts/go-async-sample/internal/usecase/interactor"
)

type CSVController struct {
	interactor interactor.CSVInteractor
}

func NewCSVController(interactor interactor.CSVInteractor) *CSVController {
	return &CSVController{interactor: interactor}
}

func (c *CSVController) HandleCSVGeneration(ctx *gin.Context) {
	if err := c.interactor.RequestCsvGenerate(ctx); err != nil {
		log.Printf("CSV generation error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process CSV generation request"})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{"message": "CSV generation request accepted"})
}
