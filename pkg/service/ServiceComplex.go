package service

import (
	"sport/pkg/repository"
	"sport/sportclub"
)

type ComplexService struct {
	repo repository.Complex
}

func NewComplexService(repo repository.Complex) *ComplexService {
	return &ComplexService{repo: repo}
}

func (c *ComplexService) GetAllComplexes() ([]sportclub.ComplexJSON, error) {
	return c.repo.GetAllComplexes()
}
