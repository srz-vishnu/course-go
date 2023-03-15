package main

import (
	"fmt"
)

func fact(number int) int {

	if number == 0 || number == 1 {

		return 1

	} else if number < 0 {

		fmt.Println("Invalid number")

		return number

	} else {

		return number * fact(number-1)

	}

}

func main() {

	//f1 := fact(0)

	//fmt.Println(f1, "\n")

	// f2 := fact(-5)

	// fmt.Println(f2, "\n")

	// f3 := fact(5)

	// fmt.Println(f3, "\n")

	f4 := fact(10)

	fmt.Println(f4, "\n")

}
