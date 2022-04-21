package main

import (
	"log"
	"sport/console"
)

func main() {
	filename := console.GetFileName()
	err := console.UploadData(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
