package schedulesRepo

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/nfnt/resize"
	"image/jpeg"
	"io"
	"log"
	"net/http"
	"os"
	"sport/pkg/repository/trainerRepo"
	"time"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetStorageRootPath() string {
	storageRootPath := os.Getenv("STORAGE_PATH")
	if storageRootPath == "" {
		storageRootPath = "./"
	}
	return storageRootPath
}
func GetStorageRootUrl() string {
	return "/assets"
}

func getImageThumbUrl(url string) (error, string) {
	fmt.Println(url)

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15")
	req.Header.Set("Host", "crossfit1905.com")
	response, e := client.Do(req)

	if e != nil {
		log.Fatal(e)
		return e, ""
	}
	defer response.Body.Close()

	//open a file for writing
	hashedFileName := GetMD5Hash(url)
	//hashedFileNameStr := string(hashedFileName[:])

	fullFilePath := GetStorageRootPath() + "/" + hashedFileName + ".jpg"
	file, err := os.Create(fullFilePath)
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	fmt.Println(response.Body)
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	fmt.Println("Success!")
	time.Sleep(3 * time.Second) // сервер блокирует слишком частые запросы

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	imageFile, err := os.Open(fullFilePath)
	img, err := jpeg.Decode(imageFile)
	if err != nil {
		fmt.Println("ERROR on decode")
		fmt.Println(err)
		return err, ""
	}

	m := resize.Resize(300, 0, img, resize.Lanczos3)
	thumbUrl := hashedFileName + "_thumb.jpg"
	fullThumbFilePath := GetStorageRootPath() + "/" + thumbUrl

	out, err := os.Create(fullThumbFilePath)
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)

	return nil, thumbUrl
}

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
	if schedule.TrainerPic != "" {
		err, thumbUrl := getImageThumbUrl(schedule.TrainerPic)
		if err == nil {
			schedule.TrainerPic = "https://crossfit-api.nihao.team" + GetStorageRootUrl() + "/" + thumbUrl
		} else {
			schedule.TrainerPic = ""
		}
	}

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
