package main

import (
	"Skyriders/controller"
	"Skyriders/repo"
	"Skyriders/routes"
	"Skyriders/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InitializeAllControllers(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database, enforcer *casbin.Enforcer) {
	SetupFlightController(router, logger, database)
	userService := SetupUserController(router, logger, database)
	SetupAuthRoutes(router, logger, userService, enforcer)
}

func SetupFlightController(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database) {
	flightRepo := repo.CreateFlightRepo(logger, database.Collection("flights"))
	flightService := service.CreateFlightService(logger, flightRepo)
	flightController := *controller.CreateFlightController(logger, flightRepo, flightService)
	flightRoutes := routes.NewFlightRoute(flightController)
	flightRoutes.FlightRoute(router)
}

func SetupUserController(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database) *service.UserService {
	userRepo := repo.CreateUserRepo(logger, database.Collection("users"))
	userService := service.CreateUserService(logger, userRepo)
	userController := *controller.CreateUserController(logger, userRepo, userService)
	userRoutes := routes.NewUserRoute(userController)
	userRoutes.UserRoute(router)
	return userService
}

func SetupAuthRoutes(router *gin.RouterGroup, logger *log.Logger, userService *service.UserService, enforcer *casbin.Enforcer) {
	authController := *controller.NewAuthController(logger, userService)
	authRoutes := routes.NewAuthRoute(authController, userService)
	authRoutes.AuthRoute(router, enforcer)
}
