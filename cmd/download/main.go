package main

import (
	"log"
	"sport/console"
)

func main() {
	urlStr := console.GetURL()
	cmplxs, err := console.DownloadComplexes(urlStr)
	if err != nil {
		log.Fatalln(err.Error())
	}
	console.PrintComplexes(cmplxs)

}
