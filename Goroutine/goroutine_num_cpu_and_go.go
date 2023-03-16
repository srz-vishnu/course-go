package main

import (
	"fmt"

	"runtime"

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

	fmt.Println(runtime.NumGoroutine()) // it prints 1 bcoz here only one on top of this println, thats go main

	go apple()

	go ball()

	fmt.Println(runtime.NumCPU())

	fmt.Println(runtime.NumGoroutine()) //3, main is go,also ball apple

	time.Sleep(time.Second)

}
