package routes

import (
	"Skyriders/controller"
	"Skyriders/middleware"
	"Skyriders/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController controller.UserController
}

func NewUserRoute(userController controller.UserController) *UserRoute {
	return &UserRoute{userController: userController}
}

func (ur *UserRoute) UserRoute(rg *gin.RouterGroup, userService *service.UserService, enforcer *casbin.Enforcer) {
	router := rg.Group("/user")
	router.GET("", ur.userController.GetAllUsers)
	router.POST("/apikey", middleware.DeserializeUser(userService), ur.userController.GenerateAPIKey)
	router.GET("/apikey", middleware.DeserializeUser(userService), ur.userController.GetAPIKey)
	router.DELETE("/apikey", middleware.DeserializeUser(userService), ur.userController.RevokeAPIKey)
}
