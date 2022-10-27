package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"user/internal/domain"
	myhttp "user/internal/http"
	"user/internal/service"
)

type RequestHandler struct {
	service *service.Service
	logger  *log.Logger
}

func NewRequestHandler(service *service.Service, logger *log.Logger) *RequestHandler {
	return &RequestHandler{
		service: service,
		logger:  logger,
	}
}

func (h *RequestHandler) Routes(r *mux.Router) {
	apiRoute := r.PathPrefix("/api/v1").Subrouter()

	apiRoute.Path("/users").Methods(http.MethodGet).HandlerFunc(myhttp.Handle(h.logger)(h.GetAllUsers))
	apiRoute.Path("/users").Methods(http.MethodPost).HandlerFunc(myhttp.Handle(h.logger)(h.Create))
	apiRoute.Path("/users").Methods(http.MethodDelete).HandlerFunc(myhttp.Handle(h.logger)(h.DeleteUser))

	apiRoute.Path("/{id:[0-9]+}").Methods(http.MethodPut).HandlerFunc(myhttp.Handle(h.logger)(h.ChangeAge))

	apiRoute.Path("/friends").Methods(http.MethodPost).HandlerFunc(myhttp.Handle(h.logger)(h.MakeFriends))
	apiRoute.Path("/friends/{id:[0-9]+}").Methods(http.MethodGet).HandlerFunc(myhttp.Handle(h.logger)(h.GetFriends))
}

func (h *RequestHandler) GetAllUsers(_ http.ResponseWriter, r *http.Request) myhttp.Response {
	var req domain.Request
	users, err := h.service.GetAllUsers(&req)
	if err != nil {
		return myhttp.InternalServerError(err)
	}
	return myhttp.OK(users)
}

func (h *RequestHandler) Create(_ http.ResponseWriter, r *http.Request) myhttp.Response {
	// bind the request
	var req domain.Request
	err := req.Bind(r)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// validate the request
	validator := domain.Validator{}
	err = validator.ValidateCreate(&req)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// check if the user already exists
	exist := h.service.CheckUser(&req)
	if exist {
		return myhttp.BadRequest(fmt.Errorf("user already exists"))
	}
	id, err := h.service.Create(&req)
	if err != nil {
		return myhttp.InternalServerError(err)
	}

	// call the service method for this
	message := fmt.Sprintf("User created with id: %d", id)
	return myhttp.OK(message)
}

func (h *RequestHandler) ChangeAge(_ http.ResponseWriter, r *http.Request) myhttp.Response {
	var req domain.Request
	err := req.Bind(r)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	err = req.BindRequestParams(r)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// validate the request
	validator := domain.Validator{}
	err = validator.ValidateChangeAge(&req)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// update the age
	err = h.service.ChangeAge(&req)
	if err != nil {
		return myhttp.InternalServerError(err)
	}
	message := fmt.Sprintf("Age updated to %d for user with id: %d", req.Age, req.TargetID)
	return myhttp.OK(message)
}

func (h *RequestHandler) GetFriends(_ http.ResponseWriter, r *http.Request) myhttp.Response {
	// bind the request
	var req domain.Request
	err := req.BindRequestParams(r)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// validate the request
	validator := domain.Validator{}
	err = validator.ValidateGetFriends(&req)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// check if the user exists
	_, err = h.service.GetUserByID(req.TargetID)
	if err != nil {
		return myhttp.BadRequest(fmt.Errorf("user does not exist"))
	}
	// get the friends
	friends, err := h.service.GetFriends(&req)
	if err != nil {
		return myhttp.InternalServerError(err)
	}
	return myhttp.OK(friends)
}

func (h *RequestHandler) MakeFriends(_ http.ResponseWriter, r *http.Request) myhttp.Response {
	// bind the request
	var req domain.Request
	err := req.Bind(r)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// validate the request
	validator := domain.Validator{}
	err = validator.ValidateMakeFriends(&req)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// check if the users exist
	// try to get the users by id, if error is not nil -> user does not exist
	_, err = h.service.GetUserByID(req.TargetID)
	if err != nil {
		return myhttp.BadRequest(fmt.Errorf("user with id %d does not exist", req.TargetID))
	}
	_, err = h.service.GetUserByID(req.SourceID)
	if err != nil {
		return myhttp.BadRequest(fmt.Errorf("user with id %d does not exist", req.SourceID))
	}
	// make friends
	err = h.service.MakeFriends(&req)
	if err != nil {
		return myhttp.InternalServerError(err)
	}
	message := fmt.Sprintf("Users with id: %d and %d are now friends", req.SourceID, req.TargetID)
	return myhttp.OK(message)
}

func (h *RequestHandler) DeleteUser(_ http.ResponseWriter, r *http.Request) myhttp.Response {
	// bind the request
	var req domain.Request
	err := req.Bind(r)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// validate the request
	validator := domain.Validator{}
	err = validator.ValidateDeleteUser(&req)
	if err != nil {
		return myhttp.BadRequest(err)
	}
	// check if the user exists
	_, err = h.service.GetUserByID(req.TargetID)
	if err != nil {
		return myhttp.BadRequest(fmt.Errorf("user does not exist"))
	}
	// delete the user
	err = h.service.DeleteUser(&req)
	if err != nil {
		return myhttp.InternalServerError(err)
	}
	message := fmt.Sprintf("User with id: %d deleted", req.TargetID)
	return myhttp.OK(message)
}
