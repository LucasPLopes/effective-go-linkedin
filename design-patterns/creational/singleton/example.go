package main

import "fmt"

func main() {
	log := getLoggerInstance()

	log.SetLevel(1)
	log.Log("This is a logger message")

	log = getLoggerInstance()
	log.SetLevel(0)
	log.Log("This is a 0 Level")

	// create several goroutines  that try to get
	// the logger instance concurrently

	for i := 1; i <= 10; i++ {
		go getLoggerInstance()
	}

	fmt.Scanln()
}
