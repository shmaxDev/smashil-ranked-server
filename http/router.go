package http

import (
	"net/http"
	"smashil-ranked/handlers"
)

func SetupRouter(mux *http.ServeMux, userHandler *handlers.UserHandler) {
	usersGroup := NewRouterGroup("/users", mux)

	usersGroup.HandleFunc("POST /", userHandler.HandlePostPlayer)
}