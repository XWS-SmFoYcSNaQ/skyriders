package routes

import (
	"Skyriders/controller"
	"Skyriders/middleware"
	"Skyriders/service"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authController controller.AuthController
	authService    *service.UserService
}

func NewAuthRoute(authController controller.AuthController, authService *service.UserService) *AuthRoute {
	return &AuthRoute{authController: authController, authService: authService}
}

func (ar *AuthRoute) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")
	router.POST("/register", ar.authController.Register)
	router.POST("/login", ar.authController.Login)
	router.GET("/refresh", ar.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(ar.authService), ar.authController.Logout)
}
