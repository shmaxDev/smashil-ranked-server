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

func (r *UserRepository) CreatePlayer(id string, discordUsername string) error {
	_, err := r.DB.Exec("INSERT INTO players (discord_id, discord_username) VALUES ($1, $2)", id, discordUsername)

	return err
}

func (r *UserRepository) GetUserById(id string) (int, error) {
	var count int

	err := r.DB.QueryRow("SELECT COUNT(*) FROM players WHERE discord_id=$1", id).Scan(&count)

	return count, err
}
