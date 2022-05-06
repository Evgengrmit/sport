package service

import (
	"sport/pkg/repository"
	"sport/pkg/service/authService"
	"sport/pkg/service/schedulesService"
	"sport/pkg/service/wodService"
)

type Service struct {
	wodService.WorkoutDay
	schedulesService.Schedule
	authService.Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{WorkoutDay: wodService.NewWorkoutDayService(repos.WorkoutDayRepo),
		Schedule:      schedulesService.NewScheduleService(repos.ScheduleRepo),
		Authorization: authService.NewAuthorizationService(repos.AuthorizationRepo)}
}
