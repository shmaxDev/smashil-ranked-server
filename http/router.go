package http

import (
	"net/http"
	"smashil-ranked/handlers"
)

func SetupRouter(mux *http.ServeMux, userHandler *handlers.UserHandler, matchHandler *handlers.MatchHandler) {
	usersGroup := NewRouterGroup("/users", mux)
	usersGroup.HandleFunc("POST /", userHandler.HandlePostPlayer)
	usersGroup.HandleFunc("POST /queue", userHandler.HandleAddToQueue)

	matchesGroup := NewRouterGroup("/matches", mux)
	matchesGroup.HandleFunc("POST /report", matchHandler.ReportMatchHandler)
}
