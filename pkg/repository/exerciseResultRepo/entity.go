package exerciseResultRepo

import "github.com/jmoiron/sqlx"

type ExerciseResult struct {
	ID         int    `json:"id"`
	UserId     int    `json:"userId"`
	ExerciseId int    `json:"exerciseId"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"createdAt,omitempty"`
}

type ExerciseResultRepo interface {
	CreateExerciseResult(ex ExerciseResult) (ExerciseResult, error)
}

type ExerciseResultRepository struct {
	db *sqlx.DB
}
