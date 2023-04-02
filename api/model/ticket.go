package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ticket struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	FlightID   primitive.ObjectID `bson:"flightId,omitempty" json:"flightId"`
	CustomerID primitive.ObjectID `bson:"customerID,omitempty" json:"customerId"`
}

type CustomerTicket struct {
	FlightId               primitive.ObjectID `bson:"flightId,omitempty" json:"flightId"`
	FlightDateSource       primitive.DateTime `bson:"flightDateSource,omitempty" json:"flightDateSource"`
	FlightDateDestination  primitive.DateTime `bson:"flightDateDestination,omitempty" json:"flightDateDestination"`
	FlightPlaceSource      string             `bson:"flightPlaceSource,omitempty" json:"flightPlaceSource"`
	FlightPlaceDestination string             `bson:"flightPlaceDestination,omitempty" json:"flightPlaceDestination"`
	FlightTicketPrice      float32            `bson:"flightTicketPrice,omitempty" json:"flightTicketPrice"`
	Quantity               int                `bson:"quantity,omitempty" json:"quantity"`
}
