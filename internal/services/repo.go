package services

import "github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"

type ISrvDb interface {
	AddUser(user domain.User) error
	DeleteUserById(id string) error
	GetAllUsers() ([]domain.User, error)
}

func (s *services) AddUser(user domain.User) error {
	err := s.data.CreateUser(s.ctx, user)
	return err
}

func (s *services) DeleteUserById(id string) error {
	err := s.data.DeleteUser(s.ctx, id)
	return err
}

func (s *services) GetAllUsers() ([]domain.User, error) {
	users, err := s.data.FindAll(s.ctx)
	return users, err
}
