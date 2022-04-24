package main

import (
	"log"
	"sport/console"
)

func main() {
	filename := console.GetFileName()
	err := console.UploadComplexes(filename)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
