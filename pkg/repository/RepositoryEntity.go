package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/schedules"
	"sport/pkg/repository/wod"
)

type Repository struct {
	wod.WorkoutDay
	schedules.Schedule
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		WorkoutDay: wod.NewWorkoutDayRepository(db),
		Schedule:   schedules.NewScheduleRepository(db),
	}
}
