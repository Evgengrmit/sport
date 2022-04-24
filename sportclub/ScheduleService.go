package sportclub

import (
	"time"
)

func (s *Schedule) GetTime() time.Time {
	timeString := s.Time
	needTime, _ := time.Parse("3:04", timeString)
	currentTime := time.Now()
	currentDay := int(currentTime.Weekday())
	needDate := currentTime.AddDate(0, 0, s.Day-currentDay)
	return time.Date(needDate.Year(), needDate.Month(), needDate.Day(), needTime.Hour(),
		needTime.Minute(), needTime.Second(), needTime.Nanosecond(), needDate.Location())
}
