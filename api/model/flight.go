package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TicketPrice      float32            `bson:"ticketPrice" json:"ticketPrice"`
	DateSource       primitive.DateTime `bson:"dateSource" json:"dateSource"`
	DateDestination  primitive.DateTime `bson:"dateDestination" json:"dateDestination"`
	PlaceSource      string             `bson:"placeSource" json:"placeSource"`
	PlaceDestination string             `bson:"placeDestination" json:"placeDestination"`
	TotalTickets     int                `bson:"totalTickets" json:"totalTickets"`
	BoughtTickets    int                `bson:"boughtTickets" json:"boughtTickets"`
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
