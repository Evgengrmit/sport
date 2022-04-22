package service

import (
	"sport/pkg/repository"
	"sport/sportclub"
)

type Service struct {
	Complex
}
type Complex interface {
	GetAllComplexes() ([]sportclub.ComplexJSON, error)
	//CreateComplex(s sportclub.SportComplex) (int, error)
	//IsComplexExists(s sportclub.SportComplex) bool
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Complex: NewComplexService(repos.Complex)}
}
