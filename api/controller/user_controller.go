package controller

import (
	"Skyriders/repo"
	"Skyriders/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get users"})
		uc.logger.Print("Unable to get users:", err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
