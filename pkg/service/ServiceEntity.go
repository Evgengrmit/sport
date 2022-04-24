package service

import (
	"sport/pkg/repository"
	"sport/sportclub"
)

type Service struct {
	Complex
	Schedule
}
type Complex interface {
	GetAllComplexes() ([]sportclub.ComplexJSON, error)
}
type Schedule interface {
	GetAllSchedules() ([]sportclub.ScheduleJSON, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Complex: NewComplexService(repos.Complex),
		Schedule: NewScheduleService(repos.Schedule)}
}
