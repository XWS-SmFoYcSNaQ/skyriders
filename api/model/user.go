package model

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"time"
)

type Role int

const (
	CustomerRole Role = iota
	AdminRole
)

type Gender int

const (
	Male Gender = iota
	Female
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Role     Role               `bson:"role" json:"role"`
	Customer *Customer          `bson:"customer,omitempty" json:"customer"`
	Admin    *Admin             `bson:"admin,omitempty" json:"admin"`
}

type Users []*User

func (p *Users) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *User) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

type Customer struct {
	Firstname   string
	Lastname    string
	DateOfBirth time.Time
	Gender      Gender
	Phone       string
	Nationality string
	//ticket slice
}

type Admin struct {
	Type string // Maybe add enum later..
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    string `json:"id,omitempty"`
	Email string `json:"email,omitempty"`
	Role  string `json:"role,omitempty"`
}
