package routes

import (
	"Skyriders/controller"
	"Skyriders/middleware"
	"Skyriders/service"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type TicketRoute struct {
	ticketController *controller.TicketController
}

func NewTicketRoute(ticketController *controller.TicketController) *TicketRoute {
	return &TicketRoute{ticketController}
}

func (ticketRoute *TicketRoute) TicketRoute(rg *gin.RouterGroup, userService *service.UserService, enforcer *casbin.Enforcer) {
	subRouter := rg.Group("/tickets")
	subRouter.POST("", middleware.AuthorizeTickets("POST", enforcer, userService),
		ticketRoute.ticketController.BuyTickets)
	subRouter.GET("", middleware.AuthorizeTickets("GET", enforcer, userService),
		ticketRoute.ticketController.GetCustomerTickets)
}
