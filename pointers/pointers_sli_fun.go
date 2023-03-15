package main

import "fmt"

func main() {

	elements := []string{"m", "n", "o"}

	slicePointer(&elements)

	fmt.Println(elements)

}

func slicePointer(s *[]string) {

	fmt.Println(*s)

	*s = append(*s, "p", "q")

}
