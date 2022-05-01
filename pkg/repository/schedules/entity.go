package schedules

import (
	"github.com/jmoiron/sqlx"
	"sport/sportclub/schedules"
)

type Schedule interface {
	GetAllSchedules() ([]schedules.ScheduleJSON, error)
	CreateSchedule(sch schedules.ScheduleJSON) (int, error)
	IsScheduleExists(sch schedules.ScheduleJSON) (bool, error)
}

type ScheduleRepository struct {
	db *sqlx.DB
}
