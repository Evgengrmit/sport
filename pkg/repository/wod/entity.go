package wod

import (
	"github.com/jmoiron/sqlx"
	"sport/sportclub/wod"
)

type WorkoutDay interface {
	GetAllWorkoutDays() ([]wod.WorkoutDayJSON, error)
	CreateWorkoutDay(s wod.WorkoutDay) (int, error)
	IsWorkoutDayExists(s wod.WorkoutDay) (bool, error)
	GetWorkoutDaysByDays() (map[string][]wod.WorkoutDayJSON, error)
}

type WorkoutDayRepository struct {
	db *sqlx.DB
}
