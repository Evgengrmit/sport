package service

import (
	"sport/pkg/repository"
	"sport/sportclub"
)

type WorkoutDayService struct {
	repo repository.WorkoutDay
}

func NewWorkoutDayService(repo repository.WorkoutDay) *WorkoutDayService {
	return &WorkoutDayService{repo: repo}
}

func (c *WorkoutDayService) GetAllWorkoutDays() ([]sportclub.ComplexJSON, error) {
	return c.repo.GetAllWorkoutDays()
}
