package console

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	club "sport/sportclub"
)

func GetSchedulesFromURL(url string) ([]club.Schedule, error) {
	body, err := GetDataFromUrl(url)
	if err != nil {
		return nil, err
	}
	var schedules []club.Schedule

	err = json.Unmarshal(body, &schedules)
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func SaveSchedulesInFile(schedules []club.Schedule) error {
	data, err := json.MarshalIndent(schedules, "", "")
	if err != nil {
		return errors.New("save schedules: " + err.Error())
	}
	err = ioutil.WriteFile("schedules.json", data, 0644)
	if err != nil {
		return errors.New("save schedules: " + err.Error())
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
