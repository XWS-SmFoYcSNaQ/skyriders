package service

import (
	"Skyriders/model"
	"Skyriders/repo"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

type UserService struct {
	logger *log.Logger
	repo   *repo.UserRepo
}

func CreateUserService(l *log.Logger, r *repo.UserRepo) *UserService {
	userService := &UserService{l, r}
	userService.addAdmin()
	return userService
}

func (service *UserService) addAdmin() {
	adminID, _ := primitive.ObjectIDFromHex("6425bd9edb1ff9554c5621da")
	user := &model.User{
		ID:       adminID,
		Email:    "admin@admin.com",
		Password: "$2a$10$0HQOLdjnsu3b1TFiP8SaG.H9ibeDQh88mZRuzBsep.ZXaN49Yqngm",
		Role:     model.AdminRole,
		Customer: nil,
		Admin: &model.Admin{
			Type: "super",
		},
	}

	_, _ = service.repo.Insert(user)
}

type CreateCustomerRequestParams struct {
	Email       string       `json:"email"`
	Password    string       `json:"password"`
	Firstname   string       `json:"firstname"`
	Lastname    string       `json:"lastname"`
	DateOfBirth time.Time    `json:"dateOfBirth"`
	Gender      model.Gender `json:"gender"`
	Phone       string       `json:"phone"`
	Nationality string       `json:"nationality"`
}

func (service *UserService) Insert(user *model.User) (id string, err error) {
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
