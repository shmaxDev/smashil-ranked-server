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

func (r *UserRepository) CreatePlayer(id string, discordUsername string) (error){
	return r.DB.QueryRow("INSERT INTO players (discord_id, discord_tag) VALUES ($1, $2)", id, discordUsername).Scan()
}

func (r *UserRepository) GetUserById(id string) (int, error){
	var count int

	err := r.DB.QueryRow("SELECT count(*) FROM players WHERE discord_id=$1", id).Scan(count)

	return count, err
}