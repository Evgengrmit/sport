package schedulesRepo

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type Schedule struct {
	ID          int       `json:"id"`
	ScheduledAt time.Time `json:"scheduledAt"`
	Name        string    `json:"name"`
	TrainerName string    `json:"trainerName"`
	TrainerPic  string    `json:"trainerPic"`
}
type ScheduleRepo interface {
	GetAllSchedules() ([]Schedule, error)
	CreateSchedule(sch Schedule) (int, error)
	IsScheduleExists(sch Schedule) (bool, error)
}

type ScheduleRepository struct {
	db *sqlx.DB
}
