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

	userRepo := repo.CreateUserRepo(logger, database.Collection("users"))
	userService := service.CreateUserService(logger, userRepo)
	userController := *controller.CreateUserController(logger, userRepo, userService)
	userRoutes := routes.NewUserRoute(userController)
	userRoutes.UserRoute(router)

	flightRepo := repo.CreateFlightRepo(logger, database.Collection("flights"))
	flightService := service.CreateFlightService(logger, flightRepo)
	flightController := *controller.CreateFlightController(logger, flightRepo, flightService)
	flightRoutes := routes.NewFlightRoute(flightController)
	flightRoutes.FlightRoute(router, userService, enforcer)

	authController := *controller.NewAuthController(logger, userService)
	authRoutes := routes.NewAuthRoute(authController)
	authRoutes.AuthRoute(router, userService, enforcer)
}
