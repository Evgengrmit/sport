package repository

import "github.com/jmoiron/sqlx"

type TrainerPostgres struct {
	db *sqlx.DB
}

func NewTrainerPostgres(db *sqlx.DB) *TrainerPostgres {
	return &TrainerPostgres{db: db}
}

func (t *TrainerPostgres) GetTrainerID(trainerName string) (int, bool) {
	exists, _ := t.IsTrainerExists(trainerName)
	if !exists {
		return 0, exists

	}
	var id int
	_ = t.db.DB.QueryRow("SELECT id FROM trainer WHERE name= $1 ",
		trainerName).Scan(&id)
	return id, exists
}
func (t *TrainerPostgres) CreateTrainer(trainerName, trainerPic string) (int, error) {
	if status, err := t.IsTrainerExists(trainerName); status || err != nil {
		return 0, err
	}
	var id int
	err := t.db.DB.QueryRow("INSERT INTO trainer (name, avatar_url) VALUES ($1, $2) RETURNING id",
		trainerName, trainerPic).Scan(&id)
	return id, err
}
func (t *TrainerPostgres) IsTrainerExists(trainerName string) (bool, error) {
	var exists bool
	err := t.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM trainer WHERE name= $1 )",
		trainerName).Scan(&exists)
	return exists, err
}
