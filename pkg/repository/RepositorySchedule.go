package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sport/sportclub"
)

type SchedulePostgres struct {
	db *sqlx.DB
}

func NewSchedulePostgres(db *sqlx.DB) *SchedulePostgres {
	return &SchedulePostgres{db: db}
}
func (s *SchedulePostgres) GetAllSchedules() ([]sportclub.ScheduleJSON, error) {
	rows, err := s.db.DB.Query("SELECT  s.id,s.name,s.scheduled_at,t.name,t.avatar_url FROM schedule s JOIN trainer t on t.id = s.trainer_id")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(rows)
	var results []sportclub.ScheduleJSON
	for rows.Next() {
		sch := sportclub.ScheduleJSON{}
		err := rows.Scan(&sch.ID, &sch.Name, &sch.ScheduledAt, &sch.TrainerName, &sch.TrainerPic)
		if err != nil {
			return nil, err
		}
		results = append(results, sch)
	}
	if err = rows.Err(); err != nil {
		return results, err
	}
	return results, nil
}
func (s *SchedulePostgres) CreateSchedule(sch sportclub.ScheduleJSON) (int, error) {
	if status, err := s.IsScheduleExists(sch); status || err != nil {
		return 0, err
	}
	trainer := NewTrainerPostgres(s.db)
	trainerID, exists := trainer.GetTrainerID(sch.TrainerName)
	var id int
	if exists {
		err := s.db.DB.QueryRow("INSERT INTO schedule (name, scheduled_at, trainer_id) VALUES ($1,$2,$3) RETURNING id",
			sch.Name, sch.ScheduledAt, trainerID).Scan(&id)
		return id, err
	}
	ctx := context.Background()
	tx, err := s.db.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	id, err = trainer.CreateTrainer(sch.TrainerName, sch.TrainerPic)
	if err != nil {
		return 0, tx.Rollback()
	}
	err = s.db.DB.QueryRow("INSERT INTO schedule (name, scheduled_at, trainer_id) VALUES ($1,$2,$3) RETURNING id",
		sch.Name, sch.ScheduledAt, trainerID).Scan(&id)
	if err != nil {
		return 0, tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *SchedulePostgres) IsScheduleExists(sch sportclub.ScheduleJSON) (bool, error) {
	var exists bool
	err := s.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM schedule WHERE name= $1 AND scheduled_at=$2)",
		sch.Name, sch.ScheduledAt).Scan(&exists)
	return exists, err

}
