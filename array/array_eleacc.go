package main

import (
	"fmt"
)

// ques to change the value 8 to 2
func main() {
	age := [4]int{1, 8, 3, 4}
	fmt.Println(age)

	age[1] = 2
	fmt.Println(age)

	//print 0 from the array
	number := [6]int{2, 4, 5, 0, 9, 74}
	fmt.Println(number[3])

}
