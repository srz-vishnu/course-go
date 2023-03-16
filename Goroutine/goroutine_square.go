package main

import (
	"fmt"

	"time"
)

func main() {

	n := 6

	go square(n)

	time.Sleep(time.Second)

}

func square(num int) int {

	result := num * num

	fmt.Println(result)

	return result

}
