package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/sportclub"
)

type Repository struct {
	Complex
}
type Complex interface {
	GetAllComplexes() ([]sportclub.ComplexJSON, error)
	CreateComplex(s sportclub.SportComplex) (int, error)
	IsComplexExists(s sportclub.SportComplex) (bool, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Complex: NewComplexPostgres(db),
	}
}
