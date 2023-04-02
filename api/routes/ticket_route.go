package routes

import (
	"Skyriders/controller"
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
	subRouter.POST("", ticketRoute.ticketController.BuyTickets) // TODO: Return middlewares
	subRouter.GET("", ticketRoute.ticketController.GetCustomerTickets)
}
