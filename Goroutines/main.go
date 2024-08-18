package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello from Goroutine")
	// Sleep for 2 seconds
	time.Sleep(time.Second * 4)
	fmt.Println("Goroutine Hello finished")
}

func sayHi() {
	fmt.Println("Hi from Goroutine")
	// Sleep for 1 second
	time.Sleep(time.Second * 1)
	fmt.Println("Goroutine Hi finished")
}

func main() {
	fmt.Println("Goroutines")
	go sayHello()
	go sayHi()

	time.Sleep(3000 * time.Millisecond)
}
