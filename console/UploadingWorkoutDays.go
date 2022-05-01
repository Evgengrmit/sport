package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sport/pkg/repository"
	"sport/sportclub/wod"
)

func GetFileName() string {
	filename := flag.String("file", "", "filename to get data from file")
	flag.Parse()
	return *filename
}

func GetWorkoutDaysFromFile(filename string) ([]wod.WorkoutDay, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var workoutDays []wod.WorkoutDay

	err = json.Unmarshal(body, &workoutDays)
	if err != nil {
		return nil, err
	}

	return workoutDays, nil
}

func AddWorkoutDaysInDB(workoutDays []wod.WorkoutDay) error {
	db, err := repository.GetConnection()
	if err != nil {
		str := fmt.Sprintf("error occurred when initialization database: %s", err.Error())
		return errors.New(str)
	}
	repos := repository.NewRepository(db)
	for _, c := range workoutDays {
		_, err = repos.CreateWorkoutDay(c)
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
