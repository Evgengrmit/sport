package schedules

import (
	"sport/pkg/repository/schedules"
	schedules2 "sport/sportclub/schedules"
)

type Schedule interface {
	GetAllSchedules() ([]schedules2.ScheduleJSON, error)
}

type ScheduleService struct {
	repo schedules.Schedule
}
