package repository

import (
	"github.com/jmoiron/sqlx"
	"sport/sportclubs"
)

type ComplexPostgres struct {
	db *sqlx.DB
}

func NewComplexPostgres(db *sqlx.DB) *ComplexPostgres {
	return &ComplexPostgres{db: db}
}
func (c *ComplexPostgres) CreateComplex(s sportclubs.SportComplex) (int, error) {
	if c.IsComplexExists(s) {
		return 0, nil
	}
	title, description, date := s.GetData()
	var id int
	err := c.db.DB.QueryRow("INSERT INTO workout_day (title, description, scheduled_at) VALUES ($1, $2, $3) RETURNING id",
		title, description, date).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (c *ComplexPostgres) IsComplexExists(s sportclubs.SportComplex) bool {
	var exists bool
	title, _, date := s.GetData()
	_ = c.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM workout_day WHERE title= $1 AND scheduled_at=$2)",
		title, date).Scan(&exists)
	return exists
}
