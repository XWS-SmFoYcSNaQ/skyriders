package repo

import (
	"Skyriders/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type TicketRepo struct {
	logger *log.Logger
	db     *mongo.Collection
}

func CreateTicketRepo(logger *log.Logger, db *mongo.Collection) *TicketRepo {
	return &TicketRepo{logger, db}
}

func (repo *TicketRepo) Insert(ticket *model.Ticket) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	result, err := repo.db.InsertOne(ctx, ticket)
	if err != nil {
		repo.logger.Println(err.Error())
		return err
	}
	repo.logger.Printf("Inserted ticket id: %v", result.InsertedID)
	return nil
}

func (repo *TicketRepo) InsertMany(tickets []model.Ticket) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	var tis []interface{}
	for _, t := range tickets {
		tis = append(tis, t)
	}
	result, err := repo.db.InsertMany(ctx, tis)
	if err != nil {
		repo.logger.Println(err.Error())
		return err
	}
	repo.logger.Printf("Inserted tickets ids: %v", result.InsertedIDs)
	return nil
}
