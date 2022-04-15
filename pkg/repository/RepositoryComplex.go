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
	if status, err := c.IsComplexExists(s); status || err != nil {
		return 0, err
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
func (c *ComplexPostgres) IsComplexExists(s sportclubs.SportComplex) (bool, error) {
	var exists bool
	title, _, date := s.GetData()
	err := c.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM workout_day WHERE title= $1 AND scheduled_at=$2)",
		title, date).Scan(&exists)
	return exists, err
}
func (c *ComplexPostgres) GetAllComplexes() ([]sportclubs.ComplexJSON, error) {
	rows, err := c.db.DB.Query("SELECT  id,title,scheduled_at FROM workout_day")
	if err != nil {
		return nil, err
	}
	results := make([]sportclubs.ComplexJSON, 0)
	for rows.Next() {
		compl := sportclubs.ComplexJSON{}
		err := rows.Scan(&compl.Id, &compl.Title, &compl.ScheduledAt)
		if err != nil {
			return nil, err
		}
		results = append(results, compl)
	}
	return results, nil
}
