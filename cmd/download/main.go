package main

import (
	"flag"
	"fmt"
	"sport/console"
)

func main() {
	urlStr := flag.String("url", "", " url for GET method")
	flag.Parse()
	err := console.GetComplexes(*urlStr)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
