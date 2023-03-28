package main

import (
	"fmt"
	"net/http"
)

// in previous exam it executes the slice value once, we need to execute it again n again, to that we use for also pass the link value to c
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

	// for i := 0; i < len(links); i++ {   // instead of for loop we can write it as below,
	// 	fmt.Println("channel msg is:", <-c)
	// }

	for {
		go Check(<-c, c)
	}
}

func Check(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "site is down")
		c <- link // we giving link to the channel
		return
	}
	fmt.Println(link, "site is up")
	c <- link
}
