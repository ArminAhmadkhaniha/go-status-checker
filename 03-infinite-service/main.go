package main

import(
	"net/http"
	"fmt"
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
	for _, link := range links{
		go checkLinkConcurrentInfinite(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLinkConcurrentInfinite(link, c)

			
		
		}(l)
	}

}

func checkLinkConcurrentInfinite(link string, c chan string) {
	_ , error := http.Get(link)
	if error != nil {
		fmt.Println(link, "is down")
		c <- link 
	} else {
		fmt.Println(link, "is up")
		c <- link
	}

}

