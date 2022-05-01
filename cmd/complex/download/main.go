package main

import (
	"log"
	"sport/console"
)

func main() {
	urlStr := console.GetURL()
	cmplxs, err := console.DownloadWorkoutDays(urlStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	console.PrintWorkoutDays(cmplxs)

}
