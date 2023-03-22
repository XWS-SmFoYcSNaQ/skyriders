package main

import (
	"Skyriders/controller"
	"Skyriders/repo"
	"Skyriders/service"
	"log"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func LoadServices(logger *log.Logger, router *mux.Router, database *mongo.Database) {
	flightRepo := repo.CreateFlightRepo(logger, database.Collection("flights"))
	userRepo := repo.CreateUserRepo(logger, database.Collection("users"))

	flightService := service.CreateFlightService(logger, flightRepo)
	userService := service.CreateUserService(logger, userRepo)

	controller.CreateFlightController(logger, router, flightRepo, flightService)
	controller.CreateUserController(logger, router, userRepo, userService)
}
