package services

import "github.com/cookienyancloud/testrpckafkapsqlclick/internal/domain"

type IQueue interface {
	SendUserLog(user domain.UserLog) error
}


func (s *services) SendUserLog(user domain.UserLog) error {
	queue := s.mq.MessageToQueue(user).
}