package main

import (
	"fmt"
	"sync"
)

// func worker(i int, wg *sync.WaitGroup) {
// 	defer wg.Done() // Signal to main goroutine that this worker has finished.  This is a critical step to prevent deadlock in the main goroutine.  Without this, the main goroutine would block forever waiting for the wg.Wait() call to return.  This is why we pass the wg pointer to the worker function.  The worker function modifies the wg counter, and the main function waits for the wg counter to reach zero before it exits.  This ensures that all
// 	fmt.Printf("Started worker %d\n", i)
// 	fmt.Printf("Ended worker %d\n", i)
// }

func Task(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(i)
}

func main() {
	// var wg sync.WaitGroup
	// for i := 1; i <= 3; i++ {
	// 	wg.Add(1)
	// 	go worker(i,&wg)
	// }
	// wg.Wait()
	// fmt.Println("ENDED MAIN FUNC")

	//Intilization of wait group
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1);
		go Task(i,&wg)
	}
	
	//End of wait group
	wg.Wait()
}
