package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"https://example.com",
		"https://google.com",
		"https://github.com",
	}

	// Channel to collect responses
	results := make(chan string)

	for _, url := range urls {
		go func(u string) {
			// Make HTTP request
			res, err := http.Get(u)
			if err != nil {
				results <- fmt.Sprintf("Error fetching %s: %s", u, err)
				return
			}
			defer res.Body.Close()
			results <- fmt.Sprintf("Fetched %s with status code: %d", u, res.StatusCode)
		}(url)
	}

	// Timeout to prevent hanging if any request takes too long
	timeout := time.After(10 * time.Second)

	// Collect responses or exit on timeout
	for i := 0; i < len(urls); i++ {
		select {
		case res := <-results:
			fmt.Println(res)
		case <-timeout:
			fmt.Println("Timed out!")
			return
		}
	}

	close(results)
}
