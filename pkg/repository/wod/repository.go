package wod

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sport/sportclub/wod"
)

func NewWorkoutDayRepository(db *sqlx.DB) *WorkoutDayRepository {
	return &WorkoutDayRepository{db: db}
}
func (c *WorkoutDayRepository) CreateWorkoutDay(s wod.WorkoutDay) (int, error) {
	if status, err := c.IsWorkoutDayExists(s); status || err != nil {
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
func (c *WorkoutDayRepository) IsWorkoutDayExists(s wod.WorkoutDay) (bool, error) {
	var exists bool
	title, _, date := s.GetData()
	err := c.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM workout_day WHERE title= $1 AND scheduled_at=$2)",
		title, date).Scan(&exists)
	return exists, err
}
func (c *WorkoutDayRepository) GetAllWorkoutDays() ([]wod.WorkoutDayJSON, error) {
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
	var results []wod.WorkoutDayJSON
	for rows.Next() {
		compl := wod.WorkoutDayJSON{}
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
func (c *WorkoutDayRepository) GetWorkoutDaysByDays() (map[string][]wod.WorkoutDayJSON, error) {
	return nil, nil
}
