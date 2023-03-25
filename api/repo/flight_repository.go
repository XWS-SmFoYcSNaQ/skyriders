package repo

import (
	"Skyriders/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepo struct {
	logger *log.Logger
	db     *mongo.Collection
}

func CreateFlightRepo(l *log.Logger, c *mongo.Collection) *FlightRepo {
	return &FlightRepo{l, c}
}

func (fr *FlightRepo) GetAll() (model.Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var flights model.Flights
	flightsCursor, err := fr.db.Find(ctx, bson.M{})
	if err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	if err = flightsCursor.All(ctx, &flights); err != nil {
		fr.logger.Println(err)
		return nil, err
	}
	return flights, nil
}

func (fr *FlightRepo) Insert(flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := fr.db.InsertOne(ctx, &flight)
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	fr.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return nil
}

func (fr *FlightRepo) Delete(flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := fr.db.DeleteOne(ctx, &flight)
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	return nil
}
