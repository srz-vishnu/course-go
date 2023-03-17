package main

import "fmt"

// Question is there anychance for deadlock on buffered channel, buffered channel is when the capacity(size) is dcleared, if not mentioned unbuffered channel

func main() {

	c := make(chan string, 1) //answer : yes, here we set the capacity to 1 but the incoming is 3, so deadlock comes
	c <- "Peter"

	c <- "John"

	c <- "Anne"

	message := <-c

	fmt.Println(message)

	message2 := <-c

	fmt.Println(message2)

	message3 := <-c

	fmt.Println(message3)

}
