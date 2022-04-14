package main

import (
	"flag"
	"fmt"
	"log"
	"sport/console"
	"sport/pkg/db"
)

func main() {
	urlStr := flag.String("url", "", " url for GET method")
	filename := flag.String("file", "", "filename to get data from file")
	flag.Parse()
	err := console.GetComplexes(*urlStr)
	if err != nil {
		fmt.Printf(err.Error())
	}
	db.ConnectPostgresDB(db.Config{
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
	err = console.AddInDB(*filename)
	if err != nil {
		fmt.Printf(err.Error())
	}

}
