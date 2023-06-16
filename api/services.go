package main

import (
	"Skyriders/controller"
	"Skyriders/repo"
	"Skyriders/routes"
	"Skyriders/service"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeAllControllers(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database, enforcer *casbin.Enforcer) {
	userRepo := repo.CreateUserRepo(logger, database.Collection("users"))
	userService := service.CreateUserService(logger, userRepo)
	userController := *controller.CreateUserController(logger, userRepo, userService)
	userRoutes := routes.NewUserRoute(userController)
	userRoutes.UserRoute(router, userService, enforcer)

	flightRepo := repo.CreateFlightRepo(logger, database.Collection("flights"))
	flightService := service.CreateFlightService(logger, flightRepo)
	flightController := *controller.CreateFlightController(logger, flightRepo, flightService)
	flightRoutes := routes.NewFlightRoute(flightController)
	flightRoutes.FlightRoute(router, userService, enforcer)

	ticketRepo := repo.CreateTicketRepo(logger, database.Collection("tickets"))
	ticketService := service.CreateTicketService(logger, ticketRepo, flightService, userService)
	ticketController := controller.CreateTicketController(logger, ticketService)
	ticketRoutes := routes.NewTicketRoute(ticketController)
	ticketRoutes.TicketRoute(router, userService, enforcer)

	authController := *controller.NewAuthController(logger, userService)
	authRoutes := routes.NewAuthRoute(authController)
	authRoutes.AuthRoute(router, userService, enforcer)
}
