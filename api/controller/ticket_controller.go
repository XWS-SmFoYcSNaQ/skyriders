package controller

import (
	"Skyriders/contracts"
	"Skyriders/mappers"
	"Skyriders/model"
	"Skyriders/service"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TicketController struct {
	logger        *log.Logger
	ticketService *service.TicketService
}

func CreateTicketController(logger *log.Logger, ticketService *service.TicketService) *TicketController {
	return &TicketController{logger, ticketService}
}

func (controller *TicketController) BuyTickets(ctx *gin.Context) {
	user := controller.validateUser(ctx)
	if user == nil {
		return
	}

	buyTicketRequest := contracts.BuyTicketRequest{}
	err := json.NewDecoder(ctx.Request.Body).Decode(&buyTicketRequest)
	if err != nil {
		controller.logger.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, "bad request body")
		return
	}

	flightId, err := primitive.ObjectIDFromHex(buyTicketRequest.FlightId)
	if err != nil {
		controller.logger.Printf(err.Error())
		ctx.JSON(http.StatusNotFound, "flight not found")
		return
	}

	err = controller.ticketService.BuyTickets(flightId, user, buyTicketRequest.Quantity)
	if err != nil {
		controller.logger.Println(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.Status(http.StatusCreated)
}

func (controller *TicketController) GetCustomerTickets(ctx *gin.Context) {
	user := controller.validateUser(ctx)
	if user == nil {
		return
	}

	ctx.JSON(http.StatusOK, mappers.MapCustomerTicketsToResponse(user))
}

func (controller *TicketController) validateUser(ctx *gin.Context) *model.User {
	user, ok := ctx.Value("currentUser").(*model.User)
	if !ok {
		controller.logger.Println("Error casting 'currentUser' from context into type model.User")
		ctx.JSON(http.StatusInternalServerError, errors.New("an unknown error occurred"))
		return nil
	}

	if user.Customer == nil {
		controller.logger.Println("user is not customer????")
		ctx.JSON(http.StatusInternalServerError, errors.New("an unknown error occurred"))
		return nil
	}

	return user
}
