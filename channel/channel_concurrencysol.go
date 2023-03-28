package main

import (
	"fmt"
	"net/http"
)

// the pgm is excuting one by one, to do it concurrentley we introduced goroutines
func main() {
	links := []string{
		"http://google.com",
		"http://golang.org",
		"http://flipkart.com",
		"http://instagram.com",
	}

	c := make(chan string) // channel c of type string is initialized, then passing this to the function n func call

	for _, link := range links {
		go Check(link, c)

	}

	// fmt.Println("channel msg is :", <-c)
	// fmt.Println("2nd channel msg is :", <-c)  instead of using println again n again we can use forloop

	for i := 0; i < len(links); i++ {
		fmt.Println("channel msg is:", <-c)
	}

}

func Check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "site is down")
		c <- "might be the site is down"
		return
	}
	fmt.Println(link, "site is up")
	c <- "the site is up"
}
