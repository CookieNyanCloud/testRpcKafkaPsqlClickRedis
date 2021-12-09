package services

import "github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"

type ISrvCache interface {
	GetUsersCache() ([]domain.User, error)
	CacheUsers(all []domain.User) error
}

func (s *services) GetUsersCache() ([]domain.User, error) {
	users, err := s.cache.GetAll(s.ctx)
	return users, err
}

func (s *services) CacheUsers(all []domain.User) error {
	err := s.cache.CacheAll(s.ctx, all)
	return err
}
