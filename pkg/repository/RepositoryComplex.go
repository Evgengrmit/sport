package repository

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sport/sportclub"
)

type ComplexPostgres struct {
	db *sqlx.DB
}

func NewComplexPostgres(db *sqlx.DB) *ComplexPostgres {
	return &ComplexPostgres{db: db}
}
func (c *ComplexPostgres) CreateComplex(s sportclub.Complex) (int, error) {
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
func (c *ComplexPostgres) IsComplexExists(s sportclub.Complex) (bool, error) {
	var exists bool
	title, _, date := s.GetData()
	err := c.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM workout_day WHERE title= $1 AND scheduled_at=$2)",
		title, date).Scan(&exists)
	return exists, err
}
func (c *ComplexPostgres) GetAllComplexes() ([]sportclub.ComplexJSON, error) {
	rows, err := c.db.DB.Query("SELECT  id,title,scheduled_at FROM workout_day")
	if err != nil {
		return nil, err
	}
	fmt.Println("dfdfdfd")
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(rows)
	var results []sportclub.ComplexJSON
	for rows.Next() {
		compl := sportclub.ComplexJSON{}
		err := rows.Scan(&compl.Id, &compl.Title, &compl.ScheduledAt)
		if err != nil {
			return nil, err
		}
		results = append(results, compl)
	}
	if err = rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}
