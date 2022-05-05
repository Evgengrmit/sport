package console

import (
	"log"
	"time"
)

type schedule struct {
	Day         int    `json:"day"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	TrainerName string `json:"trainerName"`
	TrainerPic  string `json:"trainerPic"`
}
type workoutDay struct {
	Title       string `json:"title"`
	ScheduledAt string `json:"scheduledAt"`
	Description string `json:"description"`
}

func (s *schedule) GetTime() time.Time {
	timeString := s.Time
	needTime, _ := time.Parse("3:04", timeString)
	currentTime := time.Now()
	currentDay := int(currentTime.Weekday())
	needDate := currentTime.AddDate(0, 0, s.Day-currentDay)
	return time.Date(needDate.Year(), needDate.Month(), needDate.Day(), needTime.Hour(),
		needTime.Minute(), needTime.Second(), needTime.Nanosecond(), needDate.Location())
}

func (s *workoutDay) GetData() (string, string, time.Time) {
	date, err := time.Parse("02/01/06", s.ScheduledAt)
	if err != nil {
		log.Println(err.Error())
	}
	return s.Title, s.Description, date
}
