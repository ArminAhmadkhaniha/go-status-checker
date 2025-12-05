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

	startTime := time.Now()
	for _, link := range links{
		go checkLinkConcurrent(link, c)
	}
	for i := 0; i < len(links); i++{
		fmt.Println(<-c)
	}

	fmt.Printf("Total time taken: %v\n", time.Since(startTime))

}

func checkLinkConcurrent(link string, c chan string){
	 _, err := http.Get(link)
	 if err != nil {
		c <- link + "is down"
		
	 } else {
		c <- link + "is up"
	 }


}

// func checkLinkSequential(link string){

// 	response, err := http.Get(link)
// 	if err != nil{
// 		fmt.Println(link, "is down")
// 		return
// 	} 
// 	fmt.Println(link ,"is up")
// 	response.Body.Close()
	
	
	

// }