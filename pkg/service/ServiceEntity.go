package service

import (
	"sport/pkg/repository"
	"sport/pkg/service/authService"
	"sport/pkg/service/exerciseResultService"
	"sport/pkg/service/exerciseService"
	"sport/pkg/service/schedulesService"
	"sport/pkg/service/wodService"
)

type Service struct {
	wodService.WorkoutDay
	schedulesService.Schedule
	exerciseService.Exercise
	authService.Authorization
	exerciseResultService.ExerciseResult
}

func NewService(repos *repository.Repository) *Service {
	return &Service{WorkoutDay: wodService.NewWorkoutDayService(repos.WorkoutDayRepo),
		Schedule:       schedulesService.NewScheduleService(repos.ScheduleRepo),
		Exercise:       exerciseService.NewExerciseService(repos.ExerciseRepo),
		Authorization:  authService.NewAuthorizationService(repos.AuthorizationRepo),
		ExerciseResult: exerciseResultService.NewExerciseResultService(repos.ExerciseResultRepo),
	}
}
