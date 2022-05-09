package wodRepo

import (
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/trainerRepo"
	"time"
)

type WorkoutDay struct {
	Duration    time.Duration       `json:"-"`
	Id          int                 `json:"id"`
	Title       string              `json:"title,omitempty"`
	Description string              `json:"description,omitempty"`
	ScheduledAt time.Time           `json:"scheduledAt"`
	Trainer     trainerRepo.Trainer `json:"-"`
}

type WorkoutDayRepo interface {
	GetAllWorkoutDays() ([]WorkoutDay, error)
	CreateWorkoutDay(s WorkoutDay) (int, error)
	IsWorkoutDayExists(s WorkoutDay) (bool, error)
	GetWorkoutLatest() ([]WorkoutDay, error)
}

type WorkoutDayRepository struct {
	db *sqlx.DB
}
