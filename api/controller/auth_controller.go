package controller

import (
	config2 "Skyriders/config"
	"Skyriders/model"
	"Skyriders/service"
	"Skyriders/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	service *service.UserService
	ctx     context.Context
}

func NewAuthController(service *service.UserService, ctx context.Context) *AuthController {
	return &AuthController{service: service, ctx: ctx}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var credentials *service.CreateCustomerRequestParams

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword := utils.HashPassword(credentials.Password)

	user := &model.User{
		Email:    credentials.Email,
		Password: hashedPassword,
		Role:     model.CustomerRole,
		Customer: &model.Customer{
			Firstname:   credentials.Firstname,
			Lastname:    credentials.Lastname,
			DateOfBirth: credentials.DateOfBirth,
			Gender:      credentials.Gender,
			Phone:       credentials.Phone,
			Nationality: credentials.Nationality,
		},
		Admin: nil,
	}

	err := ac.service.Insert(user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status:": "success"})
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var credentials *model.LoginInput

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := ac.service.GetByEmail(credentials.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	if err := utils.ComparePasswords(user.Password, credentials.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})
		return
	}

	config, _ := config2.LoadConfig(".")

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60,
		"/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60,
		"/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60,
		"/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}
