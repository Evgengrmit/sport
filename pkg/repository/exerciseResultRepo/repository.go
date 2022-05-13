package exerciseResultRepo

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/authRepo"
	"sport/pkg/repository/exerciseRepo"
)

func NewExerciseResultRepository(db *sqlx.DB) *ExerciseResultRepository {
	return &ExerciseResultRepository{db: db}
}

func (e *ExerciseResultRepository) CreateExerciseResult(ex ExerciseResult) (ExerciseResult, error) {
	userDB := authRepo.NewAuthorization(e.db)
	exerciseDB := exerciseRepo.NewExerciseRepository(e.db)
	userExists, err := userDB.IsUserExistsByID(ex.UserId)
	if err != nil {
		return ExerciseResult{}, err
	}
	if !userExists {
		return ExerciseResult{}, errors.New("user with this id does not exist")
	}
	exerciseExists, err := exerciseDB.IsExerciseExistsByID(ex.ExerciseId)
	if err != nil {
		return ExerciseResult{}, err
	}
	if !exerciseExists {
		return ExerciseResult{}, errors.New("exercise with this id does not exist")
	}
	err = e.db.DB.QueryRow("INSERT INTO exercise_results (exercise_id, user_id, comment) VALUES ($1,$2,$3) RETURNING id,created_at",
		ex.ExerciseId, ex.UserId, ex.Comment).Scan(&ex.ID, &ex.CreatedAt)
	return ex, err
}
