package main

import (
	"fmt"

	"time"
)

func apple() {

	var i int

	for i = 1; i <= 5; i++ {

		time.Sleep(time.Millisecond * 10)

		fmt.Println(i, " apple")

	}

}

func ball() {

	var i int

	for i = 1; i <= 5; i++ {

		time.Sleep(time.Millisecond * 10)

		fmt.Println(i, " ball")

	}

}

func main() {

	go func() { // this is anonymous function, function without name is anony func

		apple()

	}() //this is how we call the function

	go func() {

		ball()

	}()

	time.Sleep(time.Second)

}
