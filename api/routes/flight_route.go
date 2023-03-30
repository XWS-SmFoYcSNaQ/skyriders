package routes

import (
	"Skyriders/controller"
	"Skyriders/middleware"
	"Skyriders/service"
	"github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
)

type FlightRoute struct {
	flightController controller.FlightController
}

func NewFlightRoute(flightController controller.FlightController) *FlightRoute {
	return &FlightRoute{flightController: flightController}
}

func (fr *FlightRoute) FlightRoute(rg *gin.RouterGroup, authService *service.UserService, enforcer *casbin.Enforcer) {
	router := rg.Group("/flight")
	router.GET("", fr.flightController.GetAllFlights)
	router.POST("", middleware.DeserializeUser(authService),
		middleware.Authorize("flight", "POST", enforcer),
		middleware.DeserializeFlight(), fr.flightController.PostFlight)
	router.DELETE("/:id", middleware.DeserializeUser(authService),
		middleware.Authorize("flight", "DELETE", enforcer),
		fr.flightController.DeleteFlight)
}
