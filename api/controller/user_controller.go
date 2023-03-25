package controller

import (
	"Skyriders/repo"
	"Skyriders/service"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type UserController struct {
	logger  *log.Logger
	repo    *repo.UserRepo
	service *service.UserService
	ctx     context.Context
}

func CreateUserController(logger *log.Logger, repo *repo.UserRepo, service *service.UserService, ctx context.Context) *UserController {
	return &UserController{logger: logger, repo: repo, service: service, ctx: ctx}
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
		uc.logger.Fatal("Unable to get users:", err)
		return
	}

	ctx.JSON(http.StatusOK, users)
}
