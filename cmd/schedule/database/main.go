package main

import (
	"log"
	"sport/console"
)

func main() {
	filename := console.GetFileName()
	err := console.UploadSchedules(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
