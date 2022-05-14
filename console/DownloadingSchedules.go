package console

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func GetSchedulesFromURL(url string) ([]ParsedScheduleItem, error) {
	body, err := GetDataFromUrl(url)
	if err != nil {
		return nil, err
	}
	var schedules []ParsedScheduleItem

	err = json.Unmarshal(body, &schedules)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func SaveSchedulesInFile(schedules []ParsedScheduleItem) error {
	data, err := json.MarshalIndent(schedules, "", "")
	if err != nil {
		return errors.New("save schedulesRepo: " + err.Error())
	}
	err = ioutil.WriteFile("schedulesRepo.json", data, 0644)
	if err != nil {
		return errors.New("save schedulesRepo: " + err.Error())
	}
	return nil
}

func DownloadSchedules(url string) error {
	schdls, err := GetSchedulesFromURL(url)
	if err != nil {
		return errors.New("download data: " + err.Error())
	}
	err = SaveSchedulesInFile(schdls)
	if err != nil {
		return errors.New("download data: " + err.Error())
	}
	return nil
}
