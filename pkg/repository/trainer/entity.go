package trainer

import "github.com/jmoiron/sqlx"

type Trainer interface {
	GetTrainerID(trainerName string) (int, bool)
	CreateTrainer(trainerName, trainerPic string) (int, error)
	IsTrainerExists(trainerName string) (bool, error)
}

type TrainerRepository struct {
	db *sqlx.DB
}
