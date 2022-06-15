package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sport/pkg/repository"
	"sport/pkg/repository/wodRepo"
)

func GetFileName() string {
	filename := flag.String("file", "", "filename to get data from file")
	flag.Parse()
	return *filename
}

func GetWorkoutDaysFromFile(filename string) ([]ParsedWorkoutDay, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var workoutDays []ParsedWorkoutDay

	err = json.Unmarshal(body, &workoutDays)
	if err != nil {
		return nil, err
	}

	return workoutDays, nil
}

func AddWorkoutDaysInDB(workoutDays []ParsedWorkoutDay) error {
	db, err := repository.GetConnection()
	if err != nil {
		str := fmt.Sprintf("error occurred when initialization database: %s", err.Error())
		return errors.New(str)
	}
	repos := repository.NewRepository(db)
	for _, c := range workoutDays {
		workoutDate, errorParseWorkoutDate := c.GetWorkoutDate()
		if errorParseWorkoutDate != nil {
			log.Println(errorParseWorkoutDate)
			continue
		}
		workoutDay := wodRepo.WorkoutDay{Title: c.Title,
			Description: c.Description,
			ScheduledAt: workoutDate}
		_, err = repos.CreateWorkoutDay(workoutDay)
		if err != nil {
			return errors.New("add workoutDays: " + err.Error())
		}
	}
	return nil
}

func UploadWorkoutDays(filename string) error {
	workoutDaysFromFile, err := GetWorkoutDaysFromFile(filename)
	if err != nil {
		return errors.New("upload data: " + err.Error())
	}
	err = AddWorkoutDaysInDB(workoutDaysFromFile)
	if err != nil {
		return errors.New("upload data: " + err.Error())
	}
	return nil
}
