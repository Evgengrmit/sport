package main

import (
	"flag"
	"fmt"
	"sport/console"
)

func main() {

	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	filename := flag.String("file", "", "filename to get data from file")
	flag.Parse()

	err := console.AddInDB(*filename)
	if err != nil {
		fmt.Printf(err.Error())
	}

}
