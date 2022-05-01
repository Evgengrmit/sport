package console

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sport/pkg/repository"
	"sport/sportclub/schedules"
)

func GetSchedulesFromFile(filename string) ([]schedules.ScheduleJSON, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var schdls []schedules.Schedule

	err = json.Unmarshal(body, &schdls)
	if err != nil {
		return nil, err
	}
	rightSchedules := make([]schedules.ScheduleJSON, 0, len(schdls))
	for _, schedule := range schdls {
		rightSchedules = append(rightSchedules, schedules.ScheduleJSON{Name: schedule.Name, ScheduledAt: schedule.GetTime(), TrainerName: schedule.TrainerName, TrainerPic: schedule.TrainerPic})
	}

	return rightSchedules, nil
}
func AddScheduleInDB(schedules []schedules.ScheduleJSON) error {
	db, err := repository.GetConnection()
	if err != nil {
		str := fmt.Sprintf("error occurred when initialization database: %s", err.Error())
		return errors.New(str)
	}
	repos := repository.NewRepository(db)
	for _, s := range schedules {
		_, err = repos.CreateSchedule(s)
		if err != nil {
			return errors.New("add schedules: " + err.Error())
		}
	}
	return nil
}
func UploadSchedules(filename string) error {
	schdls, err := GetSchedulesFromFile(filename)
	if err != nil {
		return errors.New("upload data: " + err.Error())
	}
	err = AddScheduleInDB(schdls)
	if err != nil {
		return errors.New("upload data: " + err.Error())
	}
	return nil
}
