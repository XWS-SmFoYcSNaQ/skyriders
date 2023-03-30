package contracts

type BuyTicketRequest struct {
	FlightId   string
	Quantity   int
	CustomerId string
}
