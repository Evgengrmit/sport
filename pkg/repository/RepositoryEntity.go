package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/authRepo"
	"sport/pkg/repository/schedulesRepo"
	"sport/pkg/repository/wodRepo"
)

type Repository struct {
	authRepo.AuthorizationRepo
	wodRepo.WorkoutDayRepo
	schedulesRepo.ScheduleRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		WorkoutDayRepo:    wodRepo.NewWorkoutDayRepository(db),
		ScheduleRepo:      schedulesRepo.NewScheduleRepository(db),
		AuthorizationRepo: authRepo.NewAuthorization(db),
	}
}
