package contracts

type CustomerTicketResponse struct {
	FlightId               string  `json:"flightId"`
	FlightDateSource       string  `json:"flightDateSource"`
	FlightDateDestination  string  `json:"flightDateDestination"`
	FlightPlaceSource      string  `json:"flightPlaceSource"`
	FlightPlaceDestination string  `json:"flightPlaceDestination"`
	FlightTicketPrice      float32 `json:"flightTicketPrice"`
	Quantity               int     `json:"quantity"`
	FullName               string  `json:"fullName"`
}
