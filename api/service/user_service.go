package service

import (
	"Skyriders/model"
	"Skyriders/repo"
	"Skyriders/utils"
	"errors"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	logger *log.Logger
	repo   *repo.UserRepo
}

func CreateUserService(l *log.Logger, r *repo.UserRepo) *UserService {
	userService := &UserService{l, r}
	_, err := userService.addAdmin()
	if err != nil {
		l.Println(err.Error())
	}
	return userService
}

func (service *UserService) addAdmin() (bool, error) {
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

	user, err := service.repo.GetById(user.ID.Hex())
	if err != nil {
		return false, err
	}
	if user != nil {
		_, err = service.repo.Insert(user)
		return true, nil
	}
	return false, errors.New("failed to add admin")
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

func (service *UserService) Update(userId primitive.ObjectID, user model.User) error {
	return service.repo.Update(userId, user)
}

func (service *UserService) EmailExists(email string) bool {
	user, _ := service.repo.GetByEmail(email)
	if user != nil {
		return true
	}
	return false
}

func (service *UserService) GenerateAPIKey(expiration *time.Time) (utils.APIKey, error) {
	keyLength := 32
	letterBytes := "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	keyBytes := make([]byte, keyLength)
	for i := range keyBytes {
		keyBytes[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	apiKeyString := string(keyBytes)
	apiKey := utils.APIKey{
		KeyString:  apiKeyString,
		Expiration: expiration,
	}

	if expiration.IsZero() {
		apiKey.Expiration = nil
	}

	return apiKey, nil
}

func (service *UserService) AuthorizeAPIKey(apiKey string) (*model.User, bool) {
	user, err := service.repo.GetByAPI(apiKey)

	if err != nil {
		return nil, false
	}
	return user, true
}
