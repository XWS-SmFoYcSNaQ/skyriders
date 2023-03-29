package controller

import (
	"Skyriders/model"
	"Skyriders/repo"
	"Skyriders/service"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type KeyProduct struct{}

type FlightController struct {
	logger  *log.Logger
	repo    *repo.FlightRepo
	service *service.FlightService
	ctx     context.Context
}

func CreateFlightController(logger *log.Logger, repo *repo.FlightRepo, service *service.FlightService, ctx context.Context) *FlightController {
	return &FlightController{logger: logger, repo: repo, service: service, ctx: ctx}
}

func (fc *FlightController) GetAllFlights(ctx *gin.Context) {
	flights, err := fc.repo.GetAll()

	if err != nil {
		fc.logger.Fatal("Unable to get flights:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get flights"})
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

	flightObj := model.Flight{ID: id}
	err = fc.repo.Delete(&flightObj)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{"status:": "success"})
}
