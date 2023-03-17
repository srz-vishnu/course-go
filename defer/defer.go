package main

import (
	"fmt"
	"time"
)

func count(a string) {

	var i int

	for i = 1; i <= 5; i++ {

		time.Sleep(time.Millisecond * 10)

		fmt.Println(i, a)

	}

}

func main() { // normally the apple prints first and then the ball, just bcoz we used defer it gonna work in a diff way

	defer count("apple") //bcoz of defer it only works last untill all other functions or statement executed
	count("ball")

}
