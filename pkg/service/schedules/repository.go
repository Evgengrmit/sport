package schedules

import (
	"sport/pkg/repository/schedules"
	schedules2 "sport/sportclub/schedules"
)

func NewScheduleService(repo schedules.Schedule) *ScheduleService {
	return &ScheduleService{repo: repo}
}

func (s *ScheduleService) GetAllSchedules() ([]schedules2.ScheduleJSON, error) {
	return s.repo.GetAllSchedules()
}
