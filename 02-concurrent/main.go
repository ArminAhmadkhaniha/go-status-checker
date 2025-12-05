package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)
	var wg sync.WaitGroup

	startTime := time.Now()

	for _, link := range links {
		wg.Add(1)
		// Passing the pointer to WaitGroup
		go checkLink(link, c, &wg)
	}

	// Closer goroutine
	go func() {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}

	fmt.Printf("Concurrent check took: %v\n", time.Since(startTime))
}

func checkLink(link string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	_, err := http.Get(link)
	if err != nil {
		c <- link + " is down"
		return
	}
	c <- link + " is up"
}