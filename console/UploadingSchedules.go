package console

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sport/pkg/repository"
	"sport/pkg/repository/schedulesRepo"
)

func GetSchedulesFromFile(filename string) ([]schedulesRepo.Schedule, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}
	var schedules []ParsedScheduleItem

	err = json.Unmarshal(body, &schedules)

	if err != nil {
		return nil, err
	}
	rightSchedules := make([]schedulesRepo.Schedule, 0, len(schedules))
	for _, schedule := range schedules {
		rightSchedules = append(rightSchedules,
			schedulesRepo.Schedule{
				Name:        schedule.Name,
				ScheduledAt: schedule.GetTime(),
				TrainerName: schedule.TrainerName,
				TrainerPic:  schedule.TrainerPic})
	}

	return rightSchedules, nil
}
func AddScheduleInDB(schedules []schedulesRepo.Schedule) error {
	db, err := repository.GetConnection()
	if err != nil {
		str := fmt.Sprintf("error occurred when initialization database: %s", err.Error())
		return errors.New(str)
	}
	repos := repository.NewRepository(db)
	for _, s := range schedules {
		_, err = repos.CreateSchedule(s)
		if err != nil {
			return errors.New("add schedulesRepo: " + err.Error())
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
