package main

import (
	"log"
	"sport/console"
)

func main() {
	urlStr := console.GetURL()
	cmplxs, err := console.DownloadData(urlStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	console.PrintComplexes(cmplxs)

}
