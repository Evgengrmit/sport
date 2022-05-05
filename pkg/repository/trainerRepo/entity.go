package trainerRepo

import "github.com/jmoiron/sqlx"

type TrainerRepo interface {
	GetTrainerID(trainerName string) (int, bool)
	CreateTrainer(trainerName, trainerPic string) (int, error)
	IsTrainerExists(trainerName string) (bool, error)
}

type TrainerRepository struct {
	db *sqlx.DB
}

type Trainer struct {
	Avatar string `json:"avatar"`
	Name   string `json:"name"`
}
