package schedulesService

import (
	"sport/pkg/repository/schedulesRepo"
)

type Schedule interface {
	GetAllSchedules() ([]schedulesRepo.Schedule, error)
}

type ScheduleService struct {
	repo schedulesRepo.ScheduleRepo
}
