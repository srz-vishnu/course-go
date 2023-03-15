package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4}

	new_copy := make([]int, 3)
	fmt.Println(new_copy)
	copy(new_copy, num) //destination and then source
	fmt.Println(new_copy)
}
