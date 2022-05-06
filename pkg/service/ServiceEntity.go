package service

import (
	"sport/pkg/repository"
	"sport/pkg/service/exerciseService"
	"sport/pkg/service/schedulesService"
	"sport/pkg/service/wodService"
)

type Service struct {
	wodService.WorkoutDay
	schedulesService.Schedule
	exerciseService.Exercise
}

func NewService(repos *repository.Repository) *Service {
	return &Service{WorkoutDay: wodService.NewWorkoutDayService(repos.WorkoutDayRepo),
		Schedule: schedulesService.NewScheduleService(repos.ScheduleRepo),
		Exercise: exerciseService.NewExerciseService(repos.ExerciseRepo)}
}
