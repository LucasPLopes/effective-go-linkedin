package main

import "fmt"

// This program panics because there is no goroutine outside of `main`
// interacting with the `ch` channel

// fatal error: all goroutines are sleep  - deadlock

func main() {
	ch := make(chan int)

	ch <- 10

	v := <-ch

	fmt.Println("received ", v)
}
