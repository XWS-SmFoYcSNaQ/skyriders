package service

import (
	"Skyriders/model"
	"Skyriders/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type UserService struct {
	logger *log.Logger
	repo   *repo.UserRepo
}

func CreateUserService(l *log.Logger, r *repo.UserRepo) *UserService {
	return &UserService{l, r}
}

type CreateCustomerRequestParams struct {
	Email       string             `json:"email"`
	Password    string             `json:"password"`
	Firstname   string             `json:"firstname"`
	Lastname    string             `json:"lastname"`
	DateOfBirth primitive.DateTime `json:"dateOfBirth"`
	Gender      model.Gender       `json:"gender"`
	Phone       string             `json:"phone"`
	Nationality string             `json:"nationality"`
}

func (service *UserService) Insert(user *model.User) error {
	return service.repo.Insert(user)
}

func (service *UserService) GetByEmail(email string) (*model.User, error) {
	return service.repo.GetByEmail(email)
}

func (service *UserService) GetById(id string) (*model.User, error) {
	return service.repo.GetById(id)
}

func (service *UserService) GetAll() (model.Users, error) {
	return service.repo.GetAll()
}

func (service *UserService) Update(userId primitive.ObjectID, user model.User) error {
	return service.repo.Update(userId, user)
}
