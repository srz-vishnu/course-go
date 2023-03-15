package main

import (
	"fmt"
)

func main() {
	num := []int{1, 2, 3, 4}

	fmt.Println("slice before appending:", num)
	fmt.Println("length before:", len(num))

	num = append(num, 5, 6, 7)

	fmt.Println("slice after appending:", num)
	fmt.Println("length after:", len(num))

}
