package exerciseRepo

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewExerciseRepository(db *sqlx.DB) *ExerciseRepository {
	return &ExerciseRepository{db: db}
}

func (e *ExerciseRepository) GetAllExercises() ([]Exercise, error) {
	rows, err := e.db.DB.Query("SELECT * FROM exercise")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(rows)
	var results []Exercise
	for rows.Next() {
		ex := Exercise{}
		err := rows.Scan(&ex.ID, &ex.Name, &ex.CreatedAt)
		if err != nil {
			return nil, err
		}
		results = append(results, ex)
	}
	if err = rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}
