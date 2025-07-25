package services

import (
	"fmt"
	"net/http"
	"smashil-ranked/errors"
	"smashil-ranked/repositories"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) AddUser(discordId string, username string) error {

	if doesUserExist, err := s.doesUserExist(discordId); err != nil {
		return err
	} else if doesUserExist {
		return errors.NewHttpError(http.StatusConflict, fmt.Sprintf("User with id %s already exists", discordId), nil)
	}

	return s.repo.CreatePlayer(discordId, username)
}

func (s *UserService) doesUserExist(discordId string) (bool, error) {
	count, err := s.repo.GetUserById(discordId)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
