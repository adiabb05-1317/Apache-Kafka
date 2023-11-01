package main

import (
	"fmt"
	"sync"
)

// synchronizing is about coordinating and managing the relationships between concurrent operations to ensure they work correctly together.
// Asynchronous operations are about doing tasks in the background, allowing for parallelism and not having to wait for tasks to complete before moving on.
// calling go routines is asynchronous
// go routines are lightweight threads
// channels are used to communicate between go routines
// channels are used to synchronize go routines

func main() {
	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}
	//    wait group says adding 2 goRoutines
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("go routine 1")
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println("go routine 2")
		ch <- 3
		ch <- 4
		wg.Done()
	}(myCh, wg)

	//    wait until all go routines complete their task
	wg.Wait()
}
