package main

import (
	"fmt"
	"time"
)

func apple() {

	var i int

	for i = 1; i <= 5; i++ {

		fmt.Println(i, " apple")

	}

}

func ball() {

	var i int

	for i = 1; i <= 5; i++ {

		fmt.Println(i, " ball")

	}

}

func main() {

	go apple() //goroutine, when the term go is added before an function then its go_function

	ball() //normal function

	// in this pgm only ball is printed unless we put time.sleep, bcoz main and go function are independant,main watches only ball() to solve this issue, timer will be set
	time.Sleep(time.Second) //timer set to solve issue and print apple also
}
