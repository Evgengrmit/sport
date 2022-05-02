package wod

import (
	"sport/pkg/repository/wod"
	wod2 "sport/sportclub/wod"
)

func NewWorkoutDayService(repo wod.WorkoutDay) *WorkoutDayService {
	return &WorkoutDayService{repo: repo}
}

func (c *WorkoutDayService) GetAllWorkoutDays() ([]wod2.WorkoutDayJSON, error) {
	return c.repo.GetAllWorkoutDays()
}
func (c *WorkoutDayService) GetWorkoutDaysByDays() (map[string][]wod2.WorkoutDayJSON, error) {
	return c.repo.GetWorkoutDaysByDays()
}
