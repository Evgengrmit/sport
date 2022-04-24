package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/sportclub"
)

type Repository struct {
	Complex
	Schedule
}
type Complex interface {
	GetAllComplexes() ([]sportclub.ComplexJSON, error)
	CreateComplex(s sportclub.Complex) (int, error)
	IsComplexExists(s sportclub.Complex) (bool, error)
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
		Complex:  NewComplexPostgres(db),
		Schedule: NewSchedulePostgres(db),
	}
}
