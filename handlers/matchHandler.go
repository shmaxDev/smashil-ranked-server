package handlers

import (
	"net/http"
	"smashil-ranked/dtos"
	"smashil-ranked/services"
)

type MatchHandler struct {
	MatchService *services.MatchService
}

func NewMatchHandler(s *services.MatchService) *MatchHandler {
	return &MatchHandler{s}
}

func (h *MatchHandler) ReportMatchHandler(w http.ResponseWriter, r *http.Request) {
	var match dtos.ReportMatchDto

	err := DecodeAndValidate(&match, &w, r)
	if err != nil {
		return
	}

	if err = h.MatchService.ReportMatch(match); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Match reported"))
}
