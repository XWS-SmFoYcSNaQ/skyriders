package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DateSource       primitive.DateTime `bson:"dateSource,omitempty" json:"dateSource"`
	DateDestination  primitive.DateTime `bson:"dateDestination,omitempty" json:"dateDestination"`
	PlaceSource      string             `bson:"placeSource,omitempty" json:"placeSource"`
	PlaceDestination string             `bson:"placeDestination,omitempty" json:"placeDestination"`
	TicketPrice      float32            `bson:"ticketPrice,omitempty" json:"ticketPrice"`
	TotalTickets     int                `bson:"totalTickets,omitempty" json:"totalTickets"`
	BoughtTickets    int                `bson:"boughtTickets,omitempty" json:"boughtTickets"`
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

func (p *Flight) BuyTickets(quantity int) error {
	if quantity < 1 {
		return errors.New("quantity must be greater then 0")
	}
	remainingTickets := p.TotalTickets - p.BoughtTickets
	if remainingTickets-quantity < 0 {
		msg := fmt.Sprintf("You cannot buy %d tickets, there are only %d tickets left", quantity, remainingTickets)
		return errors.New(msg)
	}
	p.BoughtTickets += quantity

	return nil
}
