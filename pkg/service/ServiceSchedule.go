package service

import (
	"sport/pkg/repository"
	"sport/sportclub"
)

type ScheduleService struct {
	repo repository.Schedule
}

func NewScheduleService(repo repository.Schedule) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) GetAllSchedules() ([]sportclub.ScheduleJSON, error) {
	return s.repo.GetAllSchedules()
}
