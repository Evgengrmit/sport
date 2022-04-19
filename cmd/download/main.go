package main

import (
	"fmt"
	"sport/console"
	"sync"
)

func main() {
	urlStr := console.GetURL()
	cmplxs, err := console.GetComplexesFromURL(urlStr)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go console.PrintComplexes(cmplxs, wg)
	err = console.SaveComplexesInFile(cmplxs)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	wg.Wait()
}
