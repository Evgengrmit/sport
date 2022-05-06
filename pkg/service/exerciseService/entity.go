package exerciseService

import "sport/pkg/repository/exerciseRepo"

type Exercise interface {
	GetAllExercises() ([]exerciseRepo.Exercise, error)
}

type ExerciseService struct {
	repo exerciseRepo.ExerciseRepo
}
