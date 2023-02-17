package main

import "fmt"

type observer interface {
	onUpdate(data string)
}

type DataListener struct {
	Name string
}

func (dl *DataListener) onUpdate(data string) {
	fmt.Printf("Listener: %s\tgot data change: %s\n", dl.Name, data)
}
