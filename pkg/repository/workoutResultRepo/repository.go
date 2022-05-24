package workoutResultRepo

import (
	"database/sql"
	"errors"
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
