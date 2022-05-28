package trainerRepo

import (
	"github.com/jmoiron/sqlx"
	imageUtils "sport/pkg/utils"
)

func NewTrainerRepository(db *sqlx.DB) *TrainerRepository {
	return &TrainerRepository{db: db}
}

func (t *TrainerRepository) GetTrainerID(trainerName string) (int, bool) {
	exists, _ := t.IsTrainerExists(trainerName)
	if !exists {
		return 0, exists

	}
	var id int
	_ = t.db.DB.QueryRow("SELECT id FROM trainer WHERE name= $1 ",
		trainerName).Scan(&id)
	return id, exists
}
func (t *TrainerRepository) CreateTrainer(trainerName, trainerPic string) (int, error) {
	if status, err := t.IsTrainerExists(trainerName); status || err != nil {
		return 0, err
	}

	// если для тренера передана аватарка - пробуем изменить её размер
	if trainerPic != "" {
		err, thumbUrl := imageUtils.GetAvatarThumbUrl(trainerPic)
		if err == nil {
			trainerPic = imageUtils.GetStorageRootUrl() + thumbUrl
		} else {
			trainerPic = ""
		}
	}

	var id int
	err := t.db.DB.QueryRow("INSERT INTO trainer (name, avatar_url) VALUES ($1, $2) RETURNING id",
		trainerName, trainerPic).Scan(&id)
	return id, err
}
func (t *TrainerRepository) IsTrainerExists(trainerName string) (bool, error) {
	var exists bool
	err := t.db.DB.QueryRow("SELECT EXISTS(SELECT * FROM trainer WHERE name= $1 )",
		trainerName).Scan(&exists)
	return exists, err
}
