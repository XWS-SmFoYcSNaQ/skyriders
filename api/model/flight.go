package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TicketPrice      float32            `bson:"ticketPrice,omitempty" json:"ticketPrice"`
	DateSource       primitive.DateTime `bson:"dateSource,omitempty" json:"dateSource"`
	DateDestination  primitive.DateTime `bson:"dateDestination,omitempty" json:"dateDestination"`
	PlaceSource      string             `bson:"placeSource,omitempty" json:"placeSource"`
	PlaceDestination string             `bson:"placeDestination,omitempty" json:"placeDestination"`
	TotalTickets     int                `bson:"totalTickets,omitempty" json:"totalTickets"`
	BoughtTickets    int                `bson:"buyTickets,omitempty" json:"boughtTickets"`
}

type Flights []*Flight

func (p *Flights) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Flight) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Flight) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}
