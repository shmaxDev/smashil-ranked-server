package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"smashil-ranked/dtos"
	internalErrors "smashil-ranked/errors"
	"smashil-ranked/services"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(u *services.UserService) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) HandlePostPlayer(w http.ResponseWriter, r *http.Request){
	var user dtos.UserDto

	fmt.Println(r.Body)

	err := json.NewDecoder(r.Body).Decode(&user)

	fmt.Println(err, user)

	if err != nil {
		http.Error(w, "Bad request: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserService.AddUser(user.Id, user.Tag)

	if err != nil {
		var httpErr *internalErrors.HTTPError

		if ok := errors.As(err, &httpErr); ok {
			http.Error(w, httpErr.Error(), httpErr.StatusCode)
		}else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

	w.WriteHeader(http.StatusCreated)
}