package console

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type ParsedScheduleItem struct {
	Day         int    `json:"day"`
	Time        string `json:"time"`
	Name        string `json:"name"`
	TrainerName string `json:"trainerName"`
	TrainerPic  string `json:"trainerPic"`
}
type ParsedWorkoutDay struct {
	Title       string `json:"title"`
	ScheduledAt string `json:"scheduledAt"`
	Description string `json:"description"`
}

func (s *ParsedScheduleItem) GetTime() time.Time {
	timeString := s.Time

	timePart := strings.Split(timeString, ":")
	hour, _ := strconv.Atoi(timePart[0])
	minutes, _ := strconv.Atoi(timePart[1])

	currentTime := time.Now()
	currentDay := int(currentTime.Weekday())
	needDate := currentTime.AddDate(0, 0, s.Day-currentDay)
	return time.Date(needDate.Year(), needDate.Month(), needDate.Day(), hour, minutes,
		0, 0, time.FixedZone("Europe/Moscow", 3600*3)).UTC()
}

func (s *ParsedWorkoutDay) GetData() (string, string, time.Time) {
	date, err := time.Parse("02/01/06", s.ScheduledAt)
	if err != nil {
		log.Println(err.Error())
	}
	return s.Title, s.Description, date
}
