package mappers

import (
	"Skyriders/contracts"
	"Skyriders/model"
)

func MapCustomerTicketToResponse(ticket *model.CustomerTicket) *contracts.CustomerTicketResponse {
	dateFormat := "2001-01-01 15:15"
	return &contracts.CustomerTicketResponse{
		FlightId:               ticket.FlightId.Hex(),
		FlightDateSource:       ticket.FlightDateDestination.Time().Format(dateFormat),
		FlightDateDestination:  ticket.FlightDateDestination.Time().Format(dateFormat),
		FlightPlaceSource:      ticket.FlightPlaceSource,
		FlightPlaceDestination: ticket.FlightPlaceDestination,
		FlightTicketPrice:      ticket.FlightTicketPrice,
		Quantity:               ticket.Quantity,
	}
}

func MapCustomerTicketsToResponse(tickets []model.CustomerTicket) []*contracts.CustomerTicketResponse {
	customerTicketsResponse := make([]*contracts.CustomerTicketResponse, len(tickets))
	for i, ticket := range tickets {
		customerTicketsResponse[i] = MapCustomerTicketToResponse(&ticket)
	}

	return customerTicketsResponse
}
