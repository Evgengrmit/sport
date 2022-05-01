package wod

import (
	"sport/pkg/repository/wod"
	wod2 "sport/sportclub/wod"
)

type WorkoutDay interface {
	GetAllWorkoutDays() ([]wod2.WorkoutDayJSON, error)
	GetWorkoutDaysByDays() (map[string][]wod2.WorkoutDayJSON, error)
}

type WorkoutDayService struct {
	repo wod.WorkoutDay
}
