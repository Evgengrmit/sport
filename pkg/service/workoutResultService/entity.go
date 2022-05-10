package workoutResultService

import "sport/pkg/repository/workoutResultRepo"

type WorkoutResult interface {
	CreateWorkoutResult(wod workoutResultRepo.WorkoutResult) (workoutResultRepo.WorkoutResult, error)
}

type WorkoutResultService struct {
	repo workoutResultRepo.WorkoutResultRepo
}
