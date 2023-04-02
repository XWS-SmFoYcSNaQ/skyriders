package repo

import (
	"Skyriders/model"
	"Skyriders/utils"
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/mongo"
)

type FlightRepo struct {
	logger *log.Logger
	db     *mongo.Collection
}

func CreateFlightRepo(l *log.Logger, c *mongo.Collection) *FlightRepo {
	return &FlightRepo{l, c}
}

func (fr *FlightRepo) GetAll(filters map[string][]string) (model.Flights, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var flights model.Flights
	flightsCursor, err := fr.db.Find(ctx, utils.ConvertFlightFilterData(filters))
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

func (fr *FlightRepo) Create(flight *model.Flight) error {
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

func (fr *FlightRepo) GetOne(flightId primitive.ObjectID) (*model.Flight, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var flight model.Flight
	result := fr.db.FindOne(ctx, bson.M{"_id": flightId}).Decode(&flight)
	if result != nil {
		fr.logger.Println("Flight repo: error getting flight, id: ", flightId)
		return nil, result
	}

	return &flight, nil
}

func (fr *FlightRepo) Update(flightId primitive.ObjectID, flight *model.Flight) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	result, err := fr.db.UpdateByID(ctx, flightId, bson.M{"$set": flight})
	if err != nil {
		fr.logger.Println(err.Error())
		return errors.New("error updating flight")
	}
	if result.MatchedCount == 0 {
		fr.logger.Printf("There is no document with id: %s", flightId.String())
		return errors.New("invalid document id")
	}
	return nil
}

func (fr *FlightRepo) Delete(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := fr.db.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		fr.logger.Println(err)
		return err
	}
	return nil
}
