package main

import (
	"fmt"
)

// ques to print numbers  even numbers from 1 to 10 without 4
func main() {

	for i := 0; i <= 10; i += 2 {
		if i == 4 {
			continue
		}
		fmt.Println("numbers are:", i)
	}
}
