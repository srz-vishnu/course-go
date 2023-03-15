package main

import (
	"fmt"
)

func main() {

	var a int = 10
	var b int = 25
	var c int = 12

	maximum := max(a, b, c)
	fmt.Println("the max value is", maximum)

}

func max(num1, num2, num3 int) int {
	var result int

	if num1 > num2 && num1 > num3 {
		result = num1
	} else if num2 > num3 {
		result = num2
	} else {
		result = num3
	}
	return result
}
