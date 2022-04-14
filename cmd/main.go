package main

import (
	"fmt"
	"sport/console"
)

func main() {
	err := console.GetComplexes(console.GetURL())
	if err != nil {
		fmt.Printf(err.Error())
	}
}
