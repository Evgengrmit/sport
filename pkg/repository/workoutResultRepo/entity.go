package workoutResultRepo

import "github.com/jmoiron/sqlx"

type WorkoutResult struct {
	ID         int    `json:"id"`
	UserId     int    `json:"userId"`
	WorkoutId  int    `json:"workoutId"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"createdAt,omitempty"`
	TimeSecond int    `json:"timeSecond"`
	TimeCap    int    `json:"timeCap"`
}

type WorkoutResultRepo interface {
	CreateWorkoutResult(wod WorkoutResult) (WorkoutResult, error)
}

type WorkoutResultRepository struct {
	db *sqlx.DB
}
