package controller

import (
	"Skyriders/model"
	"Skyriders/repo"
	"Skyriders/service"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"

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
	filters := ctx.Request.URL.Query()
	flights, err := fc.repo.GetAll(filters)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get flights"})
		fc.logger.Println("Unable to get flights:", err)
	}

	ctx.JSON(http.StatusOK, flights)
}

func (fc *FlightController) PostFlight(ctx *gin.Context) {
	flight, exists := ctx.Get("flight")
	if !exists {
		fc.logger.Println(exists)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	flightObj, ok := flight.(*model.Flight)
	if !ok {
		fc.logger.Println(ok)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err := fc.repo.Create(flightObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"status:": "success"})
}

func (fc *FlightController) DeleteFlight(ctx *gin.Context) {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	err = fc.repo.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status:": "success"})
}
