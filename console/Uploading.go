package console

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sport/pkg/repository"
	club "sport/sportclubs"
)

func GetFileName() string {
	filename := flag.String("file", "", "filename to get data from file")
	flag.Parse()
	return *filename
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

func AddComplexInDB(complexes []club.SportComplex) error {
	db, err := repository.GetConnection()
	if err != nil {
		str := fmt.Sprintf("error occurred when initialization database: %s", err.Error())
		return errors.New(str)
	}
	repos := repository.NewRepository(db)
	for _, c := range complexes {
		_, err = repos.CreateComplex(c)
		if err != nil {
			return errors.New("add complexes: " + err.Error())
		}
	}
	return nil
}

func UploadData(filename string) error {
	cmplxs, err := GetComplexFromFile(filename)
	if err != nil {
		return errors.New("upload data: " + err.Error())
	}
	err = AddComplexInDB(cmplxs)
	if err != nil {
		return errors.New("upload data: " + err.Error())
	}
	return nil
}
