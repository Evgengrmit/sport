package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/sportclub"
)

type Repository struct {
	WorkoutDay
	Schedule
}
type WorkoutDay interface {
	GetAllWorkoutDays() ([]sportclub.ComplexJSON, error)
	CreateWorkoutDay(s sportclub.Complex) (int, error)
	IsWorkoutDayExists(s sportclub.Complex) (bool, error)
}

type Schedule interface {
	GetAllSchedules() ([]sportclub.ScheduleJSON, error)
	CreateSchedule(sch sportclub.ScheduleJSON) (int, error)
	IsScheduleExists(sch sportclub.ScheduleJSON) (bool, error)
}

type Trainer interface {
	GetTrainerID(trainerName string) (int, bool)
	CreateTrainer(trainerName, trainerPic string) (int, error)
	IsTrainerExists(trainerName string) (bool, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		WorkoutDay: NewWorkoutDayRepository(db),
		Schedule:   NewSchedulePostgres(db),
	}
}
