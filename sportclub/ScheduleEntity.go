package sportclub

import "time"

type Schedule struct {
	Day         int    `json:"day"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	TrainerName string `json:"trainerName"`
	TrainerPic  string `json:"trainerPic"`
}

type ScheduleJSON struct {
	ID          int       `json:"id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Name        string    `json:"name"`
	TrainerName string    `json:"trainerName"`
	TrainerPic  string    `json:"trainerPic"`
}
