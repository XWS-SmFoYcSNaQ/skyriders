package controller

import (
	config2 "Skyriders/config"
	"Skyriders/model"
	"Skyriders/service"
	"Skyriders/utils"
	"Skyriders/validator"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type AuthController struct {
	logger  *log.Logger
	service *service.UserService
	ctx     context.Context
}

func NewAuthController(logger *log.Logger, service *service.UserService, ctx context.Context) *AuthController {
	return &AuthController{logger: logger, service: service, ctx: ctx}
}

func (ac *AuthController) Register(ctx *gin.Context) {
	var credentials *service.CreateCustomerRequestParams

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if !validator.IsValidEmail(credentials.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email address in invalid format"})
		return
	}

	if !validator.IsValidPhoneNumber(credentials.Phone) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid phone number"})
		return
	}

	if credentials.DateOfBirth.After(time.Now()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Date of birth should be in past"})
		return
	}

	if ac.service.IsEmailExists(credentials.Email) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
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

	ac.logger.Println(config.RefreshTokenPrivateKey)

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

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken}) //delete accessToken from the response after debugging
}

func (ac *AuthController) RefreshAccessToken(ctx *gin.Context) {

	cookie, err := ctx.Cookie("refresh_token")
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "could not refresh access token"})
		return
	}

	config, _ := config2.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	user, err := ac.service.GetById(sub.(string))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "the user belonging to this token no logger exists"})
		return
	}

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60,
		"/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60,
		"/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken}) //delete accessToken from the response after debugging
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
