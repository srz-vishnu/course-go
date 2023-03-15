package main

import (
	"fmt"
)

// ques sum of atural num btw 1 to 10
func main() {

	func() {

		sum := 0
		for i := 1; i <= 10; i++ {
			sum += i

		}
		fmt.Println(sum)
	}()

}
