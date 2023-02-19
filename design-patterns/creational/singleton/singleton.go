package main

import (
	"fmt"
	"sync"
)

type Logger struct {
	level int
}

func (l *Logger) Log(s string) {
	fmt.Println(l.level, ":", s)

}

func (l *Logger) levelToText() string {
	if l.level == 1 {
		return "WARN"
	}

	return "INFO"

}

func (l *Logger) SetLevel(level int) {
	l.level = level
}

var logger *Logger

// use the sync package  to enforce goroutine safety
var once sync.Once

func getLoggerInstance() *Logger {
	once.Do(func() {
		fmt.Println("Create Logger instance")
		logger = &Logger{}
	})

	return logger
}
