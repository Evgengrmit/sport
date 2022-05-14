package wodRepo

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type WorkoutDay struct {
	Id          int       `json:"id,string"`
	Title       string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	ScheduledAt time.Time `json:"scheduledAt"`
}

type WorkoutDayRepo interface {
	GetAllWorkoutDays() ([]WorkoutDay, error)
	CreateWorkoutDay(s WorkoutDay) (int, error)
	IsWorkoutDayExists(s WorkoutDay) (bool, error)
	GetWorkoutLatest() ([]WorkoutDay, error)
	IsWorkoutDayExistsByID(wodId int) (bool, error)
}

type WorkoutDayRepository struct {
	db *sqlx.DB
}
