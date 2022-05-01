package wod

import (
	"sport/sportclub/trainer"
	"time"
)

type WorkoutDay struct {
	Title       string `json:"title"`
	ScheduledAt string `json:"scheduledAt"`
	Description string `json:"description"`
}
type WorkoutDayJSON struct {
	Duration    time.Duration   `json:"duration"`
	Id          int             `json:"id"`
	Title       string          `json:"title"`
	ScheduledAt time.Time       `json:"scheduledAt"`
	Trainer     trainer.Trainer `json:"trainer"`
}
