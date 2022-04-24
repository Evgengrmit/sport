package sportclub

import (
	"log"
	"time"
)

func (s *Complex) GetData() (string, string, time.Time) {
	date, err := time.Parse("02/01/06", s.ScheduledAt)
	if err != nil {
		log.Println(err.Error())
	}
	return s.Title, s.Description, date
}
