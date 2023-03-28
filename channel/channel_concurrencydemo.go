package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://instagram.com",
		"http://golang.orgg",
		"http://flipkart.com",
	}

	for _, link := range links {
		Check(link)

	}

}

func Check(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println("site is down", err)
		return
	}
	fmt.Println("site is up")
}
