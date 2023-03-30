package routes

import (
	"Skyriders/controller"
	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	userController controller.UserController
}

func NewUserRoute(userController controller.UserController) *UserRoute {
	return &UserRoute{userController: userController}
}

func (ur *UserRoute) UserRoute(rg *gin.RouterGroup) {
	router := rg.Group("/user")
	router.GET("", ur.userController.GetAllUsers)
}
