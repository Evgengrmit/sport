package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sport/pkg/repository"
	club "sport/sportclubs"
	"sync"
)

func GetURL() string {
	urlStr := flag.String("url", "", " url for GET method")
	flag.Parse()
	return *urlStr
}
func GetFileName() string {
	filename := flag.String("file", "", "filename to get data from file")
	flag.Parse()
	return *filename
}
func AddComplexInDB(complexes []club.SportComplex) error {
	host := os.Getenv("HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("SSL_MODE")
	db, err := repository.ConnectPostgresDB(repository.Config{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		DBName:   dbname,
		SSLMode:  sslmode,
	})
	if err != nil {
		str := fmt.Sprintf("error occurred when initialization database: %s", err.Error())
		return errors.New(str)
	}
	repos := repository.NewRepository(db)
	for _, c := range complexes {
		_, err = repos.CreateComplex(c)
		if err != nil {
			return err
		}
	}
	return nil
}
func GetComplexFromFile(filename string) ([]club.SportComplex, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var complexes []club.SportComplex

	err = json.Unmarshal(body, &complexes)
	if err != nil {
		return nil, err
	}

	return complexes, nil
}
func PrintComplexes(complexes []club.SportComplex, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, c := range complexes {
		fmt.Printf("%s %s \n", c.Title, c.ScheduledAt)
	}
}

func SaveComplexesInFile(complexes []club.SportComplex) error {
	data, err := json.MarshalIndent(complexes, "", "")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("data.json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetComplexesFromURL(url string) ([]club.SportComplex, error) {
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

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		return nil, err
	}
	var complexes []club.SportComplex

	err = json.Unmarshal(body, &complexes)
	if err != nil {
		return nil, err
	}
	return complexes, nil
}
