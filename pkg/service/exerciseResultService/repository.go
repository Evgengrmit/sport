package exerciseResultService

import (
	"sport/pkg/repository/exerciseResultRepo"
)

func NewExerciseResultService(repo exerciseResultRepo.ExerciseResultRepo) *ExerciseResultService {
	return &ExerciseResultService{repo: repo}
}
func (e *ExerciseResultService) CreateExerciseResult(ex exerciseResultRepo.ExerciseResult) (exerciseResultRepo.ExerciseResult, error) {
	return e.repo.CreateExerciseResult(ex)
}
