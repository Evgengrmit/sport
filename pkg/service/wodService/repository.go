package wodService

import (
	"sport/pkg/repository/wodRepo"
)

func NewWorkoutDayService(repo wodRepo.WorkoutDayRepo) *WorkoutDayService {
	return &WorkoutDayService{repo: repo}
}

func (c *WorkoutDayService) GetAllWorkoutDays() ([]wodRepo.WorkoutDay, error) {
	return c.repo.GetAllWorkoutDays()
}
func (c *WorkoutDayService) GetWorkoutLatest() ([]wodRepo.WorkoutDay, error) {
	return c.repo.GetWorkoutLatest()
}
