package mappers

import (
	"Skyriders/contracts"
	"Skyriders/model"
	"fmt"
)

func MapCustomerTicketToResponse(ticket *model.CustomerTicket, user *model.User) *contracts.CustomerTicketResponse {
	dateFormat := "2001-01-01 15:15"
	return &contracts.CustomerTicketResponse{
		FlightId:               ticket.FlightId.Hex(),
		FlightDateSource:       ticket.FlightDateDestination.Time().Format(dateFormat),
		FlightDateDestination:  ticket.FlightDateDestination.Time().Format(dateFormat),
		FlightPlaceSource:      ticket.FlightPlaceSource,
		FlightPlaceDestination: ticket.FlightPlaceDestination,
		FlightTicketPrice:      ticket.FlightTicketPrice,
		Quantity:               ticket.Quantity,
		FullName:               fmt.Sprintf("%s %s", user.Customer.Firstname, user.Customer.Lastname),
	}
}

func MapCustomerTicketsToResponse(user *model.User) []*contracts.CustomerTicketResponse {
	tickets := user.Customer.Tickets
	customerTicketsResponse := make([]*contracts.CustomerTicketResponse, len(tickets))
	for i, ticket := range tickets {
		customerTicketsResponse[i] = MapCustomerTicketToResponse(&ticket, user)
	}

	return customerTicketsResponse
}
