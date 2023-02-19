package main

import "fmt"

func main() {
	go hello()
}

func hello() {
	fmt.Println("It's most likely you will never see this")
}
