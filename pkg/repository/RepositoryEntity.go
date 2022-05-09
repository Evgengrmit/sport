package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/authRepo"
	"sport/pkg/repository/exerciseRepo"
	"sport/pkg/repository/exerciseResultRepo"
	"sport/pkg/repository/schedulesRepo"
	"sport/pkg/repository/wodRepo"
)

type Repository struct {
	authRepo.AuthorizationRepo
	wodRepo.WorkoutDayRepo
	schedulesRepo.ScheduleRepo
	exerciseRepo.ExerciseRepo
	exerciseResultRepo.ExerciseResultRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		WorkoutDayRepo:     wodRepo.NewWorkoutDayRepository(db),
		ScheduleRepo:       schedulesRepo.NewScheduleRepository(db),
		ExerciseRepo:       exerciseRepo.NewExerciseRepository(db),
		AuthorizationRepo:  authRepo.NewAuthorization(db),
		ExerciseResultRepo: exerciseResultRepo.NewExerciseResultRepository(db),
	}
}
