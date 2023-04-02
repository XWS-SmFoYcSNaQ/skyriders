package contracts

type BuyTicketRequest struct {
	FlightId string `json:"flightId"`
	Quantity int    `json:"quantity"`
}
