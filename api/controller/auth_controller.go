package controller

import (
	config2 "Skyriders/config"
	"Skyriders/model"
	"Skyriders/service"
	"Skyriders/utils"
	"Skyriders/validator"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	logger  *log.Logger
	service *service.UserService
}

func NewAuthController(logger *log.Logger, service *service.UserService) *AuthController {
	return &AuthController{logger: logger, service: service}
}

func (ac *AuthController) Register(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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

		if ac.service.EmailExists(credentials.Email) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
			return
		}

		hashedPassword := utils.HashPassword(credentials.Password)
		credentials.Password = ""

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

		id, err := ac.service.Insert(user)

		if err != nil {
			ctx.JSON(http.StatusBadGateway, err.Error())
			return
		}

		_, _ = enforcer.AddGroupingPolicy(fmt.Sprint(id), "customer")

		ctx.JSON(http.StatusCreated, gin.H{"status:": "success"})
	}
}

func (ac *AuthController) Login(ctx *gin.Context) {
	var credentials *model.LoginInput

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, err := ac.service.GetByEmail(credentials.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid email or password"})
		return
	}

	if err := utils.ComparePasswords(user.Password, credentials.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid email or password"})
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

	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60,
		"/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
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

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

func (ac *AuthController) Logout(ctx *gin.Context) {
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (ac *AuthController) CheckAuth(ctx *gin.Context) {
	var accessToken string

	authorizationHeader := ctx.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)

	if len(fields) != 0 && fields[0] == "Bearer" {
		accessToken = fields[1]
	} else {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "unauthenticated"})
		return
	}

	config, _ := config2.LoadConfig(".")
	_, err := utils.ValidateToken(accessToken, config.AccessTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "unauthenticated"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "authenticated"})
}
