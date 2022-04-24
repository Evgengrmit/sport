package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	club "sport/sportclub"
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
func GetComplexesFromURL(url string) ([]club.Complex, error) {
	body, err := GetDataFromUrl(url)
	if err != nil {
		return nil, err
	}
	var complexes []club.Complex

	err = json.Unmarshal(body, &complexes)
	if err != nil {
		return nil, err
	}
	return complexes, nil
}

func SaveComplexesInFile(complexes []club.Complex) error {
	data, err := json.MarshalIndent(complexes, "", "")
	if err != nil {
		return errors.New("save complexes: " + err.Error())
	}
	err = ioutil.WriteFile("complexes.json", data, 0644)
	if err != nil {
		return errors.New("save complexes: " + err.Error())
	}
	return nil
}

func PrintComplexes(complexes []club.Complex) {
	for _, c := range complexes {
		fmt.Printf("%s %s \n", c.Title, c.ScheduledAt)
	}
}

func DownloadComplexes(url string) ([]club.Complex, error) {
	cmplxs, err := GetComplexesFromURL(url)
	if err != nil {
		return nil, errors.New("download data: " + err.Error())
	}
	err = SaveComplexesInFile(cmplxs)
	if err != nil {
		return nil, errors.New("download data: " + err.Error())
	}
	return cmplxs, nil
}
