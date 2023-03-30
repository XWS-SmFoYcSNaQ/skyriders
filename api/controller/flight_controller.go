package controller

import (
	"Skyriders/model"
	"Skyriders/repo"
	"Skyriders/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type KeyProduct struct{}

type FlightController struct {
	logger  *log.Logger
	repo    *repo.FlightRepo
	service *service.FlightService
}

func CreateFlightController(logger *log.Logger, repo *repo.FlightRepo, service *service.FlightService) *FlightController {
	return &FlightController{logger: logger, repo: repo, service: service}
}

func (fc *FlightController) GetAllFlights(ctx *gin.Context) {
	flights, err := fc.repo.GetAll()

	if err != nil {
		fc.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	if err := ctx.ShouldBindJSON(&flights); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get flights"})
		fc.logger.Fatal("Unable to get flights:", err)
		return
	}

	ctx.JSON(http.StatusOK, flights)
}

func (fc *FlightController) PostFlight(ctx *gin.Context) {
	flight := ctx.Value(KeyProduct{}).(*model.Flight)
	err := fc.repo.Insert(flight)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status:": "success"})
}
