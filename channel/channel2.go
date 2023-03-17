//James has an assignment in his hand. He wants to create a channel and a goroutine. The channel should receive a string from the
//goroutine six-time and send it to the main goroutine to print out six-time.

//   Solution

package main

import (
	"fmt"
)

func main() {

	c := make(chan string)

	go count("Peter is a good boy", c)

	//fmt.Println(<-c)

	for i := 1; i <= 6; i++ {

		msg := <-c

		fmt.Println(msg)

	}

}

func count(person string, ch chan string) {

	for i := 1; i <= 6; i++ {

		ch <- person

	}

}
