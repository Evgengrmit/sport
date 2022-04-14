package sportclubs

import (
	"log"
	"sport/pkg/db"
)

func (s *SportComplex) CreateNewComplex() {
	if s.IsComplexExists() {
		return
	}
	_ = db.DB.QueryRow("insert into \"workout_day\" (title, description, scheduled_at) values ($1, $2, $3)")
}
func (s *SportComplex) IsComplexExists() bool {
	var exists bool
	err := db.DB.QueryRow("select exists(select * from \"workout_day\" where title= $1 and scheduled_at=$2)", s.Title, s.ScheduledAt).Scan(&exists)
	if err != nil {
		log.Println(err.Error())
	}
	return exists
}
