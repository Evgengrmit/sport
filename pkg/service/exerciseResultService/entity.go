package exerciseResultService

import (
	"sport/pkg/repository/exerciseResultRepo"
)

type ExerciseResult interface {
	CreateExerciseResult(ex exerciseResultRepo.ExerciseResult) (exerciseResultRepo.ExerciseResult, error)
}

type ExerciseResultService struct {
	repo exerciseResultRepo.ExerciseResultRepo
}
