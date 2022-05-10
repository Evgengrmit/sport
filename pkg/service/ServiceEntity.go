package service

import (
	"sport/pkg/repository"
	"sport/pkg/service/authService"
	"sport/pkg/service/exerciseService"
	"sport/pkg/service/schedulesService"
	"sport/pkg/service/wodService"
	"sport/pkg/service/workoutResultService"
)

type Service struct {
	wodService.WorkoutDay
	schedulesService.Schedule
	exerciseService.Exercise
	authService.Authorization
	workoutResultService.WorkoutResult
}

func NewService(repos *repository.Repository) *Service {
	return &Service{WorkoutDay: wodService.NewWorkoutDayService(repos.WorkoutDayRepo),
		Schedule:      schedulesService.NewScheduleService(repos.ScheduleRepo),
		Exercise:      exerciseService.NewExerciseService(repos.ExerciseRepo),
		Authorization: authService.NewAuthorizationService(repos.AuthorizationRepo),
		WorkoutResult: workoutResultService.NewExerciseResultService(repos.WorkoutResultRepo)}
}
