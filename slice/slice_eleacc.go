package main

import (
	"fmt"
)

// ques 1) print the element at position 2 and 4
// 2) chnage the element 5,6 to 4.5

func main() {
	num := []float64{1.2, 5.6, 6.3, 8.9}

	// solution 1
	fmt.Println("slice elements are:", num)
	fmt.Println(num[1])
	fmt.Println(num[3])

	//solution 2
	num[1] = 4.5
	fmt.Println("slice elements are:", num)

}
