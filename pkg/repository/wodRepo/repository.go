package wodRepo

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func NewWorkoutDayRepository(db *sqlx.DB) *WorkoutDayRepository {
	return &WorkoutDayRepository{db: db}
}
func (c *WorkoutDayRepository) CreateWorkoutDay(s WorkoutDay) (int, error) {
	if status, err := c.IsWorkoutDayExists(s); status || err != nil {
		return 0, err
	}

	var id int
	err := c.db.DB.QueryRow("INSERT INTO workout_day (title, description, scheduled_at) VALUES ($1, $2, $3) RETURNING id",
		s.Title, s.Description, s.ScheduledAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (c *WorkoutDayRepository) IsWorkoutDayExists(s WorkoutDay) (bool, error) {
	var exists bool

	err := c.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM workout_day WHERE title= $1 AND scheduled_at=$2)",
		s.Title, s.ScheduledAt).Scan(&exists)
	return exists, err
}
func (c *WorkoutDayRepository) GetAllWorkoutDays() ([]WorkoutDay, error) {
	rows, err := c.db.DB.Query("SELECT  id,title,scheduled_at FROM workout_day")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(rows)
	var results []WorkoutDay
	for rows.Next() {
		wod := WorkoutDay{}
		err := rows.Scan(&wod.Id, &wod.Title, &wod.ScheduledAt)
		if err != nil {
			return nil, err
		}
		results = append(results, wod)
	}
	if err = rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}
func (c *WorkoutDayRepository) GetWorkoutLatest() ([]WorkoutDay, error) {
	sqlQuery := "SELECT id, title, description, scheduled_at " +
		"FROM workout_day " +
		"WHERE scheduled_at >= NOW() - INTERVAL '90 days' " +
		"ORDER BY scheduled_at"
	rows, err := c.db.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(rows)
	var results []WorkoutDay
	for rows.Next() {
		wod := WorkoutDay{}
		err := rows.Scan(&wod.Id, &wod.Title, &wod.Description, &wod.ScheduledAt)
		if err != nil {
			return nil, err
		}
		results = append(results, wod)
	}
	if err = rows.Err(); err != nil {
		return results, err
	}
	return results, nil
	//results, err := c.GetAllWorkoutDays()
	//if err != nil {
	//	return nil, err
	//}
	//workoutDaysByDay := make(map[string][]WorkoutDay)
	//for _, res := range results {
	//	weekDay := res.ScheduledAt.Weekday().String()
	//	if _, ok := workoutDaysByDay[weekDay]; ok {
	//		workoutDaysByDay[weekDay] = append(workoutDaysByDay[weekDay], res)
	//	} else {
	//		massiveOfWod := make([]WorkoutDay, 0)
	//		massiveOfWod = append(massiveOfWod, res)
	//		workoutDaysByDay[weekDay] = massiveOfWod
	//	}
	//
	//}
}
func (c *WorkoutDayRepository) IsWorkoutDayExistsByID(wodId int) (bool, error) {
	var exists bool

	err := c.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM workout_day WHERE id= $1)",
		wodId).Scan(&exists)
	return exists, err
}
