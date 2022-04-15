package main

import (
	"flag"
	"fmt"
	"sport/console"
)

func main() {
	urlStr := flag.String("url", "", " url for GET method")
	filename := flag.String("file", "", "filename to get data from file")
	flag.Parse()
	err := console.GetComplexes(*urlStr)
	if err != nil {
		fmt.Printf(err.Error())
	}

	err = console.AddInDB(*filename)
	if err != nil {
		fmt.Printf(err.Error())
	}

}
