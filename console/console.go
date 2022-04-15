package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
func AddInDB(filename string) error {
	db, err := repository.ConnectPostgresDB(repository.Config{
		Host:     "localhost",
		Port:     "5436",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	})
	if err != nil {
		log.Fatalf("error occurred when initialization database: %s", err.Error())
	}
	repos := repository.NewRepository(db)
	file, err := os.Open(filename)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	var complexes []club.SportComplex

	err = json.Unmarshal(body, &complexes)
	if err != nil {
		return err
	}
	for _, c := range complexes {
		_, err = repos.CreateComplex(c)
		if err != nil {
			fmt.Println("fssfd")
			log.Println(err.Error())
		}
	}
	//serv := service.Service{repos}
	//res, err := serv.Complex.GetAllComplexes()
	//fmt.Println(res)
	return nil
}

func GetComplexes(url string) error {
	if url == "" {
		return errors.New("empty url")
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "test-me")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {

		return err
	}
	var complexes []club.SportComplex

	err = json.Unmarshal(body, &complexes)
	if err != nil {
		return err
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for _, c := range complexes {
			fmt.Printf("%s %s \n", c.Title, c.ScheduledAt)
		}

	}()
	go func() {
		defer wg.Done()
		data, err := json.MarshalIndent(complexes, "", "")
		if err != nil {
			log.Fatalln(err)
		}
		err = ioutil.WriteFile("data.json", data, 0644)
	}()
	wg.Wait()
	return nil
}
