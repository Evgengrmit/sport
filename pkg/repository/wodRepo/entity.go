package wodRepo

import (
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/trainerRepo"
	"time"
)

type WorkoutDay struct {
	Duration    time.Duration `json:"duration"`
	Id          int           `json:"id"`
	Title       string        `json:"title"`
	Description string
	ScheduledAt time.Time           `json:"scheduledAt"`
	Trainer     trainerRepo.Trainer `json:"trainer"`
}

type WorkoutDayRepo interface {
	GetAllWorkoutDays() ([]WorkoutDay, error)
	CreateWorkoutDay(s WorkoutDay) (int, error)
	IsWorkoutDayExists(s WorkoutDay) (bool, error)
	GetWorkoutDaysByDays() (map[string][]WorkoutDay, error)
}

type WorkoutDayRepository struct {
	db *sqlx.DB
}
