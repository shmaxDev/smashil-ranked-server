package services

import (
	"fmt"
	"net/http"
	"smashil-ranked/errors"
	"smashil-ranked/queueLoop"
	"smashil-ranked/repositories"
	"time"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) AddUser(discordId string) (string, error) {

	if doesUserExist, err := s.doesUserExist(discordId); err != nil {
		return "", err
	} else if doesUserExist {
		return "", errors.NewHttpError(http.StatusConflict, fmt.Sprintf("User with id %s already exists", discordId), nil)
	}

	return s.repo.CreatePlayer(discordId)
}

func (s *UserService) GetUserElo(discordId string) (int, error) {
	elo, err := s.repo.GetUserEloById(discordId)
	if err != nil {
		return 0, err
	}
	return elo, nil
}

func (s *UserService) doesUserExist(discordId string) (bool, error) {
	count, err := s.repo.GetUserById(discordId)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *UserService) AddUserToLoop(discordId string) error {
	if doesUserExist, err := s.doesUserExist(discordId); err != nil {
		return err
	} else if !doesUserExist {
		s.AddUser(discordId)
	}
	elo, err := s.GetUserElo(discordId)
	if err != nil {
		return err
	}

	queueLoop.Add(queueLoop.Player{UserId: discordId, Elo: elo, TimeJoined: time.Now()})
	return nil
}
