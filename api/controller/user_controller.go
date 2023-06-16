package controller

import (
	"Skyriders/model"
	"Skyriders/repo"
	"Skyriders/service"
	"Skyriders/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	logger  *log.Logger
	repo    *repo.UserRepo
	service *service.UserService
}

func CreateUserController(logger *log.Logger, repo *repo.UserRepo, service *service.UserService) *UserController {
	return &UserController{logger: logger, repo: repo, service: service}
}

func (uc *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := uc.repo.GetAll()

	if err != nil {
		uc.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to get users"})
		uc.logger.Print("Unable to get users:", err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func getUser(ctx *gin.Context, uc *UserController) *model.User {
	user, ok := ctx.Value("currentUser").(*model.User)
	if !ok {
		uc.logger.Println("Error casting 'currentUser' from context into type model.User")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("an unknown error occurred"))
	}
	return user
}

func (uc *UserController) GenerateAPIKey(ctx *gin.Context) {
	user := getUser(ctx, uc)

	expiration := time.Time{}
	_ = ctx.ShouldBindJSON(&expiration)

	if !expiration.IsZero() && expiration.Before(time.Now()) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "bad API Key expiration date"})
		return
	}

	key, err := uc.service.GenerateAPIKey(&expiration)
	if err != nil {
		uc.logger.Println(err.Error())
	}

	user.Customer.APIKey = key

	err = uc.service.Update(user.ID, *user)
	if err != nil {
		uc.logger.Println("Error updating user")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("an unknown error occurred"))
		return
	}
	ctx.JSON(http.StatusCreated, key)
}

func (uc *UserController) RevokeAPIKey(ctx *gin.Context) {
	user := getUser(ctx, uc)

	if user.Customer.APIKey.KeyString != "" {
		user.Customer.APIKey = utils.APIKey{}
	}

	err := uc.service.Update(user.ID, *user)
	if err != nil {
		uc.logger.Println("Error updating user")
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("an unknown error occurred"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (uc *UserController) GetAPIKey(ctx *gin.Context) {
	user := getUser(ctx, uc)

	userWithAPI, err := uc.service.GetById(user.ID.Hex())
	if err != nil {
		uc.logger.Println("Failed to retrieve user from the database:", err.Error())
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errors.New("an unknown error occurred"))
		return
	}
	if userWithAPI.Customer.APIKey.KeyString == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "API Key is not registered yet"})
		return
	}
	if userWithAPI.Customer.APIKey.IsExpired() {
		ctx.AbortWithStatusJSON(http.StatusGone, gin.H{"error": "API Key has expired"})
		return
	}

	ctx.JSON(http.StatusOK, userWithAPI.Customer.APIKey)
}
