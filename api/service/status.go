package service

import (
	"github.com/haradayoshitsugucz/purple-server/domain/repository"
)

type StatusService interface {
	Ping() error
}

type StatusServiceImpl struct {
	statusRepo repository.StatusRepository
}

func NewStatusService(statusRepo repository.StatusRepository) StatusService {
	return &StatusServiceImpl{
		statusRepo: statusRepo,
	}
}

func (s *StatusServiceImpl) Ping() error {
	return s.statusRepo.Find()
}
