package main

import (
	"log"
	"sport/console"
)

func main() {
	urlStr := console.GetURL()
	err := console.DownloadSchedules(urlStr)
	if err != nil {
		log.Fatalln(err.Error())
	}

}
