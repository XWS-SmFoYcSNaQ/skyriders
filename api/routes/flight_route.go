package routes

import (
	"Skyriders/controller"
	"Skyriders/middleware"
	"github.com/gin-gonic/gin"
)

type FlightRoute struct {
	flightController controller.FlightController
}

func NewFlightRoute(flightController controller.FlightController) *FlightRoute {
	return &FlightRoute{flightController: flightController}
}

func (fr *FlightRoute) FlightRoute(rg *gin.RouterGroup) {
	router := rg.Group("/flight")
	router.GET("", fr.flightController.GetAllFlights) //maybe path "/"
	router.POST("", middleware.DeserializeFlight(), fr.flightController.PostFlight)
}
