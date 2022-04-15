package service

import (
	"sport/pkg/repository"
	"sport/sportclubs"
)

type Service struct {
	Complex
}
type Complex interface {
	GetAllComplexes() ([]sportclubs.ComplexJSON, error)
	//CreateComplex(s sportclubs.SportComplex) (int, error)
	//IsComplexExists(s sportclubs.SportComplex) bool
}

func NewService(repos *repository.Repository) *Service {
	return &Service{Complex: NewComplexService(repos.Complex)}
}
