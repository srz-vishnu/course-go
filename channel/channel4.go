package main

import "fmt"

// Question is to print the recived cannel inputs in order

func main() {

	c := make(chan string, 3)
	c <- "Peter"
	c <- "John"
	c <- "Anne"

	// message := <-c
	// fmt.Println(message)
	// message2 := <-c
	// fmt.Println(message2)
	// message3 := <-c
	// fmt.Println(message3)

	for message2 := range c {
		fmt.Println(message2)

	}

}
