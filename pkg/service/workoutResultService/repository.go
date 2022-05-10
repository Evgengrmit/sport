package workoutResultService

import "sport/pkg/repository/workoutResultRepo"

func NewExerciseResultService(repo workoutResultRepo.WorkoutResultRepo) *WorkoutResultService {
	return &WorkoutResultService{repo: repo}
}
func (w *WorkoutResultService) CreateWorkoutResult(wod workoutResultRepo.WorkoutResult) (workoutResultRepo.WorkoutResult, error) {
	return w.repo.CreateWorkoutResult(wod)
}
