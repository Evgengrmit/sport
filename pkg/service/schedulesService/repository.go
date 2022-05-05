package schedulesService

import (
	"sport/pkg/repository/schedulesRepo"
)

func NewScheduleService(repo schedulesRepo.ScheduleRepo) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) GetAllSchedules() ([]schedulesRepo.Schedule, error) {
	return s.repo.GetAllSchedules()
}
