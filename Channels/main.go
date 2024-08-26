package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Print(numChannel chan int){
	for num := range numChannel {
		fmt.Println("Number Processing: ", num);
		time.Sleep(time.Second )
	}
}

func main() {


	numChannel := make(chan int)
	go Print(numChannel)

	for {
		numChannel <- rand.Intn(100)
	}

	// messageChannel := make(chan string)

	// messageChannel <- "ping"

	// msg := <-messageChannel

	// fmt.Println("Received message:", msg)

}
