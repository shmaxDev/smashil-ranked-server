package repositories

import (
	"database/sql"
	"smashil-ranked/models"
)

type MatchRepository struct {
	db *sql.DB
}

func NewMatchRepository(db *sql.DB) *MatchRepository {
	return &MatchRepository{db}
}

func (r *MatchRepository) Create(match *models.Match) error {
    var winnerID string
    if match.Player1Score > match.Player2Score {
        winnerID = match.Player1ID
    } else if match.Player2Score > match.Player1Score {
        winnerID = match.Player2ID
    } else {
        winnerID = "0"
    }

    _, err := r.db.Exec(
        `INSERT INTO matches (player1_id, player2_id, winner_id, player1_score, player2_score, created_at)
         VALUES ($1, $2, $3, $4, $5, $6)`,
        match.Player1ID, match.Player2ID, winnerID, match.Player1Score, match.Player2Score, match.CreatedAt,
    )
    return err
}
