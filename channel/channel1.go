package main

import "fmt"

// pgm: a channel recives a string from goroutine and sends it outside

func count(person string, ch chan string) {

	ch <- person

}

func main() {

	var c = make(chan string)
	go count("peter is a boy", c)
	fmt.Println(<-c)

	// or instead of 17 we can print it like this

	// msg := <-c
	// fmt.Println(msg)

}
