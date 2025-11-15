package repositories

import (
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreatePlayer(id string) (string, error) {
	_, err := r.DB.Exec("INSERT INTO players (discord_id) VALUES ($1)", id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *UserRepository) GetUserById(id string) (int, error) {
	var count int

	err := r.DB.QueryRow("SELECT COUNT(*) FROM players WHERE discord_id=$1", id).Scan(&count)

	return count, err
}

func (r *UserRepository) GetUserEloById(id string) (int, error) {
	var elo int

	err := r.DB.QueryRow("SELECT elo FROM players WHERE discord_id=$1", id).Scan(&elo)

	return elo, err
}
