package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/sportclubs"
)

type Repository struct {
	Complex
}
type Complex interface {
	CreateComplex(s sportclubs.SportComplex) (int, error)
	IsComplexExists(s sportclubs.SportComplex) bool
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Complex: NewComplexPostgres(db),
	}
}
