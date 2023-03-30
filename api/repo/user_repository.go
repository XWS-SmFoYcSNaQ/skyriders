package repo

import (
	"Skyriders/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"time"
)

type UserRepo struct {
	logger *log.Logger
	db     *mongo.Collection
}

func CreateUserRepo(l *log.Logger, c *mongo.Collection) *UserRepo {
	return &UserRepo{l, c}
}

func (ur *UserRepo) GetAll() (model.Users, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var users model.Users
	usersCursor, err := ur.db.Find(ctx, bson.M{})
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	if err = usersCursor.All(ctx, &users); err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	return users, nil
}

func (ur *UserRepo) Insert(user *model.User) (id string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := ur.db.InsertOne(ctx, &user)
	if err != nil {
		ur.logger.Println(err)
		return "", err
	}
	ur.logger.Printf("Documents ID: %v\n", result.InsertedID)
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (ur *UserRepo) GetByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.getCollection()

	var user model.User
	err := usersCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepo) GetById(id string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usersCollection := ur.getCollection()

	var user model.User
	objId, _ := primitive.ObjectIDFromHex(id)
	err := usersCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
	if err != nil {
		ur.logger.Println(err)
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) getCollection() *mongo.Collection {
	patientDatabase := ur.db.Database()
	patientsCollection := patientDatabase.Collection("users")
	return patientsCollection
}
