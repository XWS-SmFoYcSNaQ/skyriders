package routes

import (
	"Skyriders/controller"
	"Skyriders/service"
	"github.com/gin-gonic/gin"
)

type TicketRoute struct {
	ticketController *controller.TicketController
}

func NewTicketRoute(ticketController *controller.TicketController) *TicketRoute {
	return &TicketRoute{ticketController}
}

func (ticketRoute *TicketRoute) TicketRoute(rg *gin.RouterGroup, userService *service.UserService) {
	subRouter := rg.Group("/tickets")
	// subRouter.Use(middleware.DeserializeUser(userService))
	subRouter.GET("", ticketRoute.ticketController.GetAllUserTickets)
	subRouter.POST("", ticketRoute.ticketController.BuyTickets)
}
