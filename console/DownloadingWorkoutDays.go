package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetURL() string {
	urlStr := flag.String("url", "", " url for GET method")
	flag.Parse()
	return *urlStr
}
func GetDataFromUrl(url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("empty url")
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "test-me")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
func GetWorkoutDaysFromURL(url string) ([]ParsedWorkoutDay, error) {
	body, err := GetDataFromUrl(url)
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

func SaveWorkoutDaysInFile(workoutDays []ParsedWorkoutDay) error {
	data, err := json.MarshalIndent(workoutDays, "", "")
	if err != nil {
		return errors.New("save workoutDays: " + err.Error())
	}
	err = ioutil.WriteFile("workoutDays.json", data, 0644)
	if err != nil {
		return errors.New("save workoutDays: " + err.Error())
	}
	return nil
}

func PrintWorkoutDays(workoutDays []ParsedWorkoutDay) {
	for _, d := range workoutDays {
		fmt.Printf("%s %s \n", d.Title, d.ScheduledAt)
	}
}

func DownloadWorkoutDays(url string) ([]ParsedWorkoutDay, error) {
	wods, err := GetWorkoutDaysFromURL(url)
	if err != nil {
		return nil, errors.New("download data: " + err.Error())
	}
	err = SaveWorkoutDaysInFile(wods)
	if err != nil {
		return nil, errors.New("download data: " + err.Error())
	}
	return wods, nil
}
