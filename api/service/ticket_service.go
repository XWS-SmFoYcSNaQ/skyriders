package service

import (
	"Skyriders/model"
	"Skyriders/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type TicketService struct {
	logger        *log.Logger
	ticketRepo    *repo.TicketRepo
	flightService *FlightService
	UserService   *UserService
}

func CreateTicketService(logger *log.Logger, ticketRepo *repo.TicketRepo, flightService *FlightService, UserService *UserService) *TicketService {
	return &TicketService{logger, ticketRepo, flightService, UserService}
}

func (ticketService *TicketService) BuyTickets(flightId primitive.ObjectID, user *model.User, quantity int) error {
	flight, err := ticketService.flightService.ReserveTickets(flightId, quantity)
	if err != nil {
		return err
	}

	var newTickets = make([]model.Ticket, quantity)
	for i := 0; i < quantity; i++ {
		newTicket := model.Ticket{FlightID: flightId, CustomerID: user.ID}
		newTickets[i] = newTicket
	}
	err = ticketService.ticketRepo.InsertMany(newTickets)
	if err != nil {
		return err
	}

	if user.Customer.Tickets == nil {
		user.Customer.Tickets = make([]model.CustomerTicket, 0, 1)
	}

	idx := getTicketIdxForFlight(user.Customer.Tickets, flightId)
	if idx != -1 {
		user.Customer.Tickets[idx].Quantity += quantity
	} else {
		customerTicket := createCustomerTicket(*flight, quantity)
		user.Customer.Tickets = append(user.Customer.Tickets, *customerTicket)
	}

	err = ticketService.UserService.Update(user.ID, *user)
	if err != nil {
		return err
	}

	return nil
}

func createCustomerTicket(flight model.Flight, quantity int) *model.CustomerTicket {

	customerTicket := &model.CustomerTicket{
		FlightId:               flight.ID,
		FlightDateSource:       flight.DateSource,
		FlightDateDestination:  flight.DateDestination,
		FlightPlaceSource:      flight.PlaceSource,
		FlightPlaceDestination: flight.PlaceDestination,
		FlightTicketPrice:      flight.TicketPrice,
		Quantity:               quantity,
	}

	return customerTicket
}

func getTicketIdxForFlight(customerTickets []model.CustomerTicket, flightId primitive.ObjectID) int {
	for i, t := range customerTickets {
		if t.FlightId == flightId {
			return i
		}
	}

	return -1
}
