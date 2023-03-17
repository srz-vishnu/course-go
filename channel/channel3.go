// Solution

package main

import (
	"fmt"

	"time"
)

func main() {

	c := make(chan string)

	go count("Peter is a good boy", c)

	//fmt.Println(<-c)

	for msg := range c {

		fmt.Println(msg)

	}

}

func count(person string, ch chan string) {

	for i := 1; i <= 6; i++ {

		ch <- person

		time.Sleep(time.Millisecond * 500)

	}

	close(ch)

}
