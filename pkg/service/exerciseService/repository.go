package exerciseService

import "sport/pkg/repository/exerciseRepo"

func NewExerciseService(repo exerciseRepo.ExerciseRepo) *ExerciseService {
	return &ExerciseService{repo: repo}
}

func (e *ExerciseService) GetAllExercises() ([]exerciseRepo.Exercise, error) {
	return e.repo.GetAllExercises()
}
