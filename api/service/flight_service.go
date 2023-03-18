package service

import (
	"Skyriders/repo"
	"log"
)

type FlightService struct {
	logger *log.Logger
	repo   *repo.FlightRepo
}

func CreateFlightService(l *log.Logger, r *repo.FlightRepo) *FlightService {
	return &FlightService{l, r}
}
