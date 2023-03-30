package main

import (
	"Skyriders/controller"
	"Skyriders/repo"
	"Skyriders/routes"
	"Skyriders/service"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeAllControllers(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database) {
	userService := SetupUserController(router, logger, database)
	SetupAuthRoutes(router, logger, userService)
	flightService := SetupFlightController(router, logger, database)
	SetupTicketController(router, logger, database, flightService, userService)
}

func SetupFlightController(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database) *service.FlightService {
	flightRepo := repo.CreateFlightRepo(logger, database.Collection("flights"))
	flightService := service.CreateFlightService(logger, flightRepo)
	flightController := *controller.CreateFlightController(logger, flightRepo, flightService)
	flightRoutes := routes.NewFlightRoute(flightController)
	flightRoutes.FlightRoute(router)
	return flightService
}

func SetupUserController(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database) *service.UserService {
	userRepo := repo.CreateUserRepo(logger, database.Collection("users"))
	userService := service.CreateUserService(logger, userRepo)
	userController := *controller.CreateUserController(logger, userRepo, userService)
	userRoutes := routes.NewUserRoute(userController)
	userRoutes.UserRoute(router)
	return userService
}

func SetupAuthRoutes(router *gin.RouterGroup, logger *log.Logger, userService *service.UserService) {
	authController := *controller.NewAuthController(logger, userService)
	authRoutes := routes.NewAuthRoute(authController, userService)
	// router.Use(middleware.DeserializeUser(userService))
	authRoutes.AuthRoute(router)
}

func SetupTicketController(router *gin.RouterGroup,
	logger *log.Logger,
	database *mongo.Database,
	flightService *service.FlightService,
	userService *service.UserService) {
	ticketRepo := repo.CreateTicketRepo(logger, database.Collection("tickets"))
	ticketService := service.CreateTicketService(logger, ticketRepo, flightService, userService)
	ticketController := controller.CreateTicketController(logger, ticketService)
	ticketRoutes := routes.NewTicketRoute(ticketController)
	ticketRoutes.TicketRoute(router, userService)
}
