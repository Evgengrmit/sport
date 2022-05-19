package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/authRepo"
	"sport/pkg/repository/exerciseRepo"
	"sport/pkg/repository/exerciseResultRepo"
	"sport/pkg/repository/schedulesRepo"
	"sport/pkg/repository/wodRepo"
	"sport/pkg/repository/workoutResultRepo"
)

type Repository struct {
	authRepo.AuthCodeRepo
	authRepo.AuthorizationRepo
	wodRepo.WorkoutDayRepo
	schedulesRepo.ScheduleRepo
	exerciseRepo.ExerciseRepo
	workoutResultRepo.WorkoutResultRepo
	exerciseResultRepo.ExerciseResultRepo
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		AuthCodeRepo:       authRepo.NewAuthCode(db),
		WorkoutDayRepo:     wodRepo.NewWorkoutDayRepository(db),
		ScheduleRepo:       schedulesRepo.NewScheduleRepository(db),
		ExerciseRepo:       exerciseRepo.NewExerciseRepository(db),
		AuthorizationRepo:  authRepo.NewAuthorization(db),
		ExerciseResultRepo: exerciseResultRepo.NewExerciseResultRepository(db),
		WorkoutResultRepo:  workoutResultRepo.NewWorkoutResultRepository(db),
	}
}
