package main

import (
	"fmt"
	"net/http"
	"time"
)

// prbl is it executing continuesly as fast as possible we need to pause in btw

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

	// for l := range c { // l is link, c for channel
	// 	go Check(l, c)   // to solve this issue
	// }

	for l := range c {
		go func() { //anonymous func
			time.Sleep(5 * time.Second)
			Check(l, c)
		}() // anonymous func call
	}
}

func Check(link string, c chan string) {
	time.Sleep(5 * time.Second)
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "site is down")
		c <- link // we giving link to the channel
		return
	}
	fmt.Println(link, "site is up")
	c <- link
}
