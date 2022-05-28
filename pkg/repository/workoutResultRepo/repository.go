package workoutResultRepo

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/authRepo"
	"sport/pkg/repository/wodRepo"
)

func NewWorkoutResultRepository(db *sqlx.DB) *WorkoutResultRepository {
	return &WorkoutResultRepository{db: db}
}

func (w *WorkoutResultRepository) CreateWorkoutResult(wod WorkoutResult) (WorkoutResult, error) {
	userDB := authRepo.NewAuthorization(w.db)
	wodDB := wodRepo.NewWorkoutDayRepository(w.db)
	userExists, err := userDB.IsUserExistsByID(wod.UserId)
	if err != nil {
		return WorkoutResult{}, err
	}
	if wod.UserId == 0 && wod.UserName == "" {
		return WorkoutResult{}, errors.New("must be a user name or id")
	}
	if !userExists && wod.UserName == "" {
		return WorkoutResult{}, errors.New("user with this id does not exist")
	}
	exerciseExists, err := wodDB.IsWorkoutDayExistsByID(wod.WorkoutId)
	if err != nil {
		return WorkoutResult{}, err
	}
	if !exerciseExists {
		return WorkoutResult{}, errors.New("workout day with this id does not exist")
	}
	err = w.db.DB.QueryRow("INSERT INTO workout_result (workout_id, user_id,user_name, comment, time_second,time_cap) VALUES ($1,$2,$3,$4,$5, $6) RETURNING id,created_at",
		wod.WorkoutId, newNullInt64(wod.UserId, userExists), newNullString(wod.UserName), wod.Comment, wod.TimeSecond, wod.TimeCap).Scan(&wod.ID, &wod.CreatedAt)
	return wod, err
}

func newNullString(name string) sql.NullString {
	if name == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: name, Valid: true}
}

func newNullInt64(id int, idExists bool) sql.NullInt64 {
	if idExists {
		return sql.NullInt64{Int64: int64(id), Valid: true}
	}
	return sql.NullInt64{}
}

func (w *WorkoutResultRepository) GetWorkoutResults(id int) ([]WorkoutResult, error) {
	rows, err := w.db.DB.Query("SELECT  id,user_name,time_second, time_cap,comment,created_at FROM workout_result WHERE workout_id=$1", id)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(rows)
	var results []WorkoutResult
	for rows.Next() {
		wod := WorkoutResult{}
		var userName sql.NullString
		err := rows.Scan(&wod.ID, &userName, &wod.TimeSecond, &wod.TimeCap, &wod.Comment, &wod.CreatedAt)
		wod.UserName = userName.String
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
