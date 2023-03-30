package service

import (
	"Skyriders/model"
	"Skyriders/repo"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FlightService struct {
	logger *log.Logger
	repo   *repo.FlightRepo
}

func CreateFlightService(l *log.Logger, r *repo.FlightRepo) *FlightService {
	return &FlightService{l, r}
}

func (service *FlightService) ReserveTickets(flightId primitive.ObjectID, quantity int) (*model.Flight, error) {
	flight, err := service.repo.GetOne(flightId)
	if err != nil {
		return nil, err
	}

	err = flight.BuyTickets(quantity)
	if err != nil {
		return nil, err
	}

	err = service.repo.Update(flightId, flight)
	if err != nil {
		return nil, err
	}

	return flight, nil
}
