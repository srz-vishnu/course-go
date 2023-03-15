package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println("odd numbers are", i)
		}

	}
}
