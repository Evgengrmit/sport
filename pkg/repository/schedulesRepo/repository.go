package schedulesRepo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"sport/pkg/repository/trainerRepo"
)

func NewScheduleRepository(db *sqlx.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}
func (s *ScheduleRepository) GetAllSchedules() ([]Schedule, error) {

	sqlQuery := "SELECT " +
		"	s.id, s.name, s.scheduled_at, t.name, t.avatar_url " +
		"FROM schedule s JOIN trainer t on t.id = s.trainer_id ORDER BY s.scheduled_at"

	rows, err := s.db.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(rows)
	var results []Schedule
	for rows.Next() {
		sch := Schedule{}
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
func (s *ScheduleRepository) CreateSchedule(schedule Schedule) (int, error) {
	if status, err := s.IsScheduleExists(schedule); status || err != nil {
		return 0, err
	}
	trainer := trainerRepo.NewTrainerRepository(s.db)

	trainerID, exists := trainer.GetTrainerID(schedule.TrainerName)
	var id int
	if exists {
		err := s.db.DB.QueryRow("INSERT INTO schedule (name, scheduled_at, trainer_id) VALUES ($1,$2,$3) RETURNING id",
			schedule.Name, schedule.ScheduledAt, trainerID).Scan(&id)
		return id, err
	}
	ctx := context.Background()
	tx, err := s.db.DB.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	id, err = trainer.CreateTrainer(schedule.TrainerName, schedule.TrainerPic)
	if err != nil {
		return 0, tx.Rollback()
	}
	err = s.db.DB.QueryRow("INSERT INTO schedule (name, scheduled_at, trainer_id) VALUES ($1,$2,$3) RETURNING id",
		schedule.Name, schedule.ScheduledAt, trainerID).Scan(&id)
	if err != nil {
		return 0, tx.Rollback()
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *ScheduleRepository) IsScheduleExists(sch Schedule) (bool, error) {
	var exists bool
	err := s.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM schedule WHERE name= $1 AND scheduled_at=$2)",
		sch.Name, sch.ScheduledAt).Scan(&exists)
	return exists, err

}
