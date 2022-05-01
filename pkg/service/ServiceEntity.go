package service

import (
	"sport/pkg/repository"
	"sport/pkg/service/schedules"
	"sport/pkg/service/wod"
)

type Service struct {
	wod.WorkoutDay
	schedules.Schedule
}

func NewService(repos *repository.Repository) *Service {
	return &Service{WorkoutDay: wod.NewWorkoutDayService(repos.WorkoutDay),
		Schedule: schedules.NewScheduleService(repos.Schedule)}
}
