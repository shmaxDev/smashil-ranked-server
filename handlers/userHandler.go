package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"smashil-ranked/dtos"
	internalErrors "smashil-ranked/errors"
	queueloop "smashil-ranked/queueLoop"
	"smashil-ranked/services"
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(u *services.UserService) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) HandlePostPlayer(w http.ResponseWriter, r *http.Request) {
	var user dtos.UserDto

	err := DecodeAndValidate(&user, &w, r)
	if err != nil {
		return
	}

	err = h.UserService.AddUser(*user.Id, *user.Username)

	if err != nil {
		var httpErr *internalErrors.HTTPError

		if ok := errors.As(err, &httpErr); ok {
			http.Error(w, httpErr.Error(), httpErr.StatusCode)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) HandleAddToQueue(w http.ResponseWriter, r *http.Request) {
	var nig dtos.Queue

	err := json.NewDecoder(r.Body).Decode(&nig)

	if err != nil {
		if errors.Is(err, io.EOF) {
			http.Error(w, "Request body is empty", http.StatusBadRequest)
			return
		}

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var player = queueloop.Player{UserId: *nig.UserId, Elo: *nig.Elo, TimeJoined: time.Now()}

	queueloop.Add(player)
}
