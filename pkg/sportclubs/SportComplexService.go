package sportclubs

import (
	"log"
	"sport/pkg/db"
	"time"
)

func (s *SportComplex) CreateNewComplex() {
	if s.IsComplexExists() {
		return
	}
	date, err := time.Parse("02/01/06", s.ScheduledAt)
	if err != nil {
		log.Println(err.Error())
	}
	_ = db.DB.QueryRow("INSERT INTO workout_day (title, description, scheduled_at) VALUES ($1, $2, $3)", s.Title, s.Description, date)
}
func (s *SportComplex) IsComplexExists() bool {
	var exists bool
	date, err := time.Parse("02/01/06", s.ScheduledAt)
	if err != nil {
		log.Println(err.Error())
	}
	err = db.DB.QueryRow("SELECT EXISTS(SELECT * FROM workout_day WHERE title= $1 AND scheduled_at=$2)", s.Title, date).Scan(&exists)
	if err != nil {
		log.Println(err.Error())
	}
	return exists
}
