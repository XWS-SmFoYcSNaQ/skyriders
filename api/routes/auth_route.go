package routes

import (
	"Skyriders/controller"
	"Skyriders/middleware"
	"Skyriders/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type AuthRoute struct {
	authController controller.AuthController
}

func NewAuthRoute(authController controller.AuthController) *AuthRoute {
	return &AuthRoute{authController: authController}
}

func (ar *AuthRoute) AuthRoute(rg *gin.RouterGroup, authService *service.UserService, enforcer *casbin.Enforcer) {
	router := rg.Group("/auth")
	router.POST("/register", middleware.Anonymous(), ar.authController.Register(enforcer))
	router.POST("/login", middleware.Anonymous(), ar.authController.Login)
	router.GET("/refresh", ar.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(authService),
		middleware.Authorize("logout", "GET", enforcer), ar.authController.Logout)
}
