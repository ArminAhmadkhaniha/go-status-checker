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
	startTime := time.Now()
	for _, link := range links{
		checkLink(link)
	}

	fmt.Printf("Total time taken: %v\n", time.Since(startTime))

}


func checkLink(link string){

	response, err := http.Get(link)
	if err != nil{
		fmt.Println(link, "is down")
		return
	} 
	fmt.Println(link ,"is up")
	response.Body.Close()
	
	
	

}