package exerciseRepo

import "github.com/jmoiron/sqlx"

type Exercise struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
type ExerciseRepo interface {
	GetAllExercises() ([]Exercise, error)
	IsExerciseExistsByID(exID int) (bool, error)
}

type ExerciseRepository struct {
	db *sqlx.DB
}
