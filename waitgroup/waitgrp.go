package main

import (
	"fmt"

	"sync"

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

		time.Sleep(time.Millisecond * 100)

		fmt.Println(i, " ball")

	}

}

func main() {

	var wg sync.WaitGroup

	wg.Add(2) // add() add tells how many goroutines are there on main goroutine

	go func() {

		apple()

		wg.Done() // Done() it tells when the excecution of the goroutine is over

	}()

	go func() {

		ball()

		wg.Done()

	}()

	wg.Wait() // Wait() we denote number of goroutines here and it reduces the value itself when one goroutine is over
}
