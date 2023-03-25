package model

import (
	"encoding/json"
	"io"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Flight struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	TicketPrice     float32            `bson:"ticketPrice,omitempty" json:"ticketPrice"`
	DateSource      primitive.DateTime `bson:"dateSource,omitempty" json:"dateSource"`
	DateDestination primitive.DateTime `bson:"dateDestination,omitempty" json:"dateDestination"`
	Pid             string             `bson:"pid,omitempty" json:"pid"`
	DateOfBirth     primitive.DateTime `bson:"dateOfBirth,omitempty" json:"dateOfBirth"`
	PhoneNumber     string             `bson:"phoneNumber,omitempty" json:"phoneNumber"`
	Nationality     string             `bson:"nationality,omitempty" json:"nationality"`
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
