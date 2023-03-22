package controller

import (
	"Skyriders/repo"
	"Skyriders/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type UserController struct {
	logger  *log.Logger
	router  *mux.Router
	repo    *repo.UserRepo
	service *service.UserService
}

func CreateUserController(l *log.Logger, r *mux.Router, repo *repo.UserRepo, s *service.UserService) *UserController {
	uc := UserController{l, r, repo, s}
	uc.registerRoutes()
	return &uc
}

func (uc *UserController) registerRoutes() {
	getRouter := uc.router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", uc.getAllUsers)
}

func (uc *UserController) getAllUsers(rw http.ResponseWriter, h *http.Request) {
	users, err := uc.repo.GetAll()

	if err != nil {
		uc.logger.Print("Database exception: ", err)
	}

	if users == nil {
		return
	}

	err = users.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		uc.logger.Fatal("Unable to convert to json :", err)
		return
	}
}
