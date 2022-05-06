package wodService

import (
	"sport/pkg/repository/wodRepo"
)

type WorkoutDay interface {
	GetAllWorkoutDays() ([]wodRepo.WorkoutDay, error)
	GetWorkoutLatest() ([]wodRepo.WorkoutDay, error)
}

type WorkoutDayService struct {
	repo wodRepo.WorkoutDayRepo
}
