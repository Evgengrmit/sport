package service

import (
	"sport/pkg/repository"
	"sport/sportclubs"
)

type ComplexService struct {
	repo repository.Complex
}

func NewComplexService(repo repository.Complex) *ComplexService {
	return &ComplexService{repo: repo}
}

func (c *ComplexService) GetAllComplexes() ([]sportclubs.ComplexJSON, error) {
	return c.repo.GetAllComplexes()
}
