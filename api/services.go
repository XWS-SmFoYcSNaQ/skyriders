package main

import (
	"Skyriders/controller"
	"Skyriders/repo"
	"Skyriders/routes"
	"Skyriders/service"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func InitializeAllControllers(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database, ctx context.Context) {
	SetupFlightController(router, logger, database, ctx)
	userService := SetupUserController(router, logger, database, ctx)
	SetupAuthRoutes(router, logger, userService)
}

func SetupFlightController(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database, ctx context.Context) {
	flightRepo := repo.CreateFlightRepo(logger, database.Collection("flights"))
	flightService := service.CreateFlightService(logger, flightRepo)
	flightController := *controller.CreateFlightController(logger, flightRepo, flightService, ctx)
	flightRoutes := routes.NewFlightRoute(flightController)
	flightRoutes.FlightRoute(router)
}

func SetupUserController(router *gin.RouterGroup, logger *log.Logger, database *mongo.Database, ctx context.Context) *service.UserService {
	userRepo := repo.CreateUserRepo(logger, database.Collection("users"))
	userService := service.CreateUserService(logger, userRepo)
	userController := *controller.CreateUserController(logger, userRepo, userService, ctx)
	userRoutes := routes.NewUserRoute(userController)
	userRoutes.UserRoute(router)
	return userService
}

func SetupAuthRoutes(router *gin.RouterGroup, logger *log.Logger, userService *service.UserService) {
	authController := *controller.NewAuthController(logger, userService, ctx)
	authRoutes := routes.NewAuthRoute(authController, userService)
	authRoutes.AuthRoute(router)
}
