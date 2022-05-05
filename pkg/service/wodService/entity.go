package wodService

import (
	"sport/pkg/repository/wodRepo"
)

type WorkoutDay interface {
	GetAllWorkoutDays() ([]wodRepo.WorkoutDay, error)
	GetWorkoutDaysByDays() (map[string][]wodRepo.WorkoutDay, error)
}

type WorkoutDayService struct {
	repo wodRepo.WorkoutDayRepo
}
