package main

import (
	"fmt"
	"sport/console"
)

func main() {
	filename := console.GetFileName()

	cmplxs, err := console.GetComplexFromFile(filename)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	err = console.AddComplexInDB(cmplxs)
	if err != nil {
		fmt.Printf(err.Error())
	}

}
