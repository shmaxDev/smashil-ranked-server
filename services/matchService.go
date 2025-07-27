package services

import (
	"smashil-ranked/dtos"
	"smashil-ranked/models"
	"smashil-ranked/repositories"
)

type MatchService struct {
	repo *repositories.MatchRepository
}

func NewMatchService(repo *repositories.MatchRepository) *MatchService {
	return &MatchService{repo}
}

func (s *MatchService) ReportMatch(m dtos.ReportMatchDto) error {
	match := &models.Match{
		Player1ID:    *m.Player1Id,
		Player2ID:    *m.Player2Id,
		Player1Score: *m.Player1Score,
		Player2Score: *m.Player2Score,
	}

	return s.repo.Create(match)
}
