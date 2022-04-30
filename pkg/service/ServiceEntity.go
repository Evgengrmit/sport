package service

import (
	"sport/pkg/repository"
	"sport/sportclub"
)

type Service struct {
	WorkoutDay
	Schedule
}
type WorkoutDay interface {
	GetAllWorkoutDays() ([]sportclub.ComplexJSON, error)
}
type Schedule interface {
	GetAllSchedules() ([]sportclub.ScheduleJSON, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{WorkoutDay: NewWorkoutDayService(repos.WorkoutDay),
		Schedule: NewScheduleService(repos.Schedule)}
}
