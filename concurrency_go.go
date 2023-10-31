package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	//  using go keword for go routine
	go greeter("hello")
	// we are  telling go to create thread to call greeter function or do sometask and we should tell when thread should come
	// back using time,sleep
	greeter("world")
	time.Sleep(100 * time.Millisecond)

	// Channels are a Go data structure that can both send and receive values.
	// They provide a way for goroutines to communicate with each other
	// and synchronize their execution.
	// Channels can be thought of as pipes that
	// connect concurrent goroutines. You can send values into channels from one goroutine and receive those values in another goroutine.
	// getting status code
	websiteList := []string{
		"https://google.com",
		"https://fb.com",
		"https://go.dev",
	}
	for _, web := range websiteList {
		getStatusCode(web)
	}
}
func greeter(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("ooops wrong endpoint or status code err")
	}
	fmt.Println("res status code:", res, endpoint)

}
