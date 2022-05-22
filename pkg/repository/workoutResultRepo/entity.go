package workoutResultRepo

import "github.com/jmoiron/sqlx"

type WorkoutResult struct {
	ID         int    `json:"id"`
	UserId     int    `json:"userId,omitempty"`
	UserName   string `json:"userName,omitempty"`
	WorkoutId  int    `json:"workoutId,omitempty"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"createdAt,omitempty"`
	TimeSecond int    `json:"timeSecond"`
	TimeCap    int    `json:"timeCap"`
}

type WorkoutResultRepo interface {
	CreateWorkoutResult(wod WorkoutResult) (WorkoutResult, error)
	GetWorkoutResults(id int) ([]WorkoutResult, error)
}

type WorkoutResultRepository struct {
	db *sqlx.DB
}
