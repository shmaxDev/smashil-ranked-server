package services

import (
	"net/http"
	"smashil-ranked/errors"
	"smashil-ranked/repositories"
	"fmt"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) AddUser(discordId string, tag string) (error){
	

	if doesUserExist, err := s.doesUserExist(discordId); err != nil {
		return err
	} else if doesUserExist {
		return errors.NewHttpError(http.StatusConflict, fmt.Sprintf("User with id %s already exists", discordId), nil)
	}


	return s.repo.CreatePlayer(discordId, tag)
}

func (s *UserService) doesUserExist(discordId string)(bool, error) {
	count, err := s.repo.GetPlayerById(discordId)

	if (err != nil){
		return false, err
	}

	return count > 0, nil
}