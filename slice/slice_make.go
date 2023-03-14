package main

//ques: slice using make

import "fmt"

func main() {

	n1 := make([]int, 2, 6)

	fmt.Printf("Slice1 = %v\n", n1)

	fmt.Printf("length_slice1 = %d\n", len(n1))

	fmt.Printf("cap_slice1 = %d\n", cap(n1))

	n2 := make([]int, 4) //here the capacity and length will be 4

	fmt.Printf("Slice2 = %v\n", n2)

	fmt.Printf("length_slice2 = %d\n", len(n2))

	fmt.Printf("cap_slice2 = %d\n", cap(n2))

}
