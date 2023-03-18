package controller

import (
	"Skyriders/model"
	"Skyriders/repo"
	"Skyriders/service"
	"context"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type FlightController struct {
	logger  *log.Logger
	router  *mux.Router
	repo    *repo.FlightRepo
	service *service.FlightService
}

func CreateFlightController(l *log.Logger, r *mux.Router, repo *repo.FlightRepo, s *service.FlightService) *FlightController {
	fc := FlightController{l, r, repo, s}
	fc.registerRoutes()
	return &fc
}

func (fc FlightController) registerRoutes() {
	getRouter := fc.router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", fc.getAllFlights)

	postRouter := fc.router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", fc.postFlight)
	postRouter.Use(fc.middlewareFlightDeserialization)
}

func (fc FlightController) getAllFlights(rw http.ResponseWriter, h *http.Request) {
	flights, err := fc.repo.GetAll()

	if err != nil {
		fc.logger.Print("Database exception: ", err)
	}

	if flights == nil {
		return
	}

	err = flights.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		fc.logger.Fatal("Unable to convert to json :", err)
		return
	}
}

func (fc *FlightController) postFlight(rw http.ResponseWriter, h *http.Request) {
	flight := h.Context().Value(KeyProduct{}).(*model.Flight)
	fc.repo.Insert(flight)
	rw.WriteHeader(http.StatusCreated)
}

func (fc *FlightController) middlewareFlightDeserialization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		flight := &model.Flight{}
		err := flight.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			fc.logger.Fatal(err)
			return
		}
		ctx := context.WithValue(h.Context(), KeyProduct{}, flight)
		h = h.WithContext(ctx)
		next.ServeHTTP(rw, h)
	})
}
