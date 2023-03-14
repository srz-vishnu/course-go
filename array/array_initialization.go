package main

import "fmt"

func main() {

	a1 := [3]int{} //without value initializtaion so it will set the default value depend on its type

	a2 := [5]int{4, 6, 8}

	a3 := [7]string{"a", "b", "c", "d", "e", "f", "g"}

	a4 := [7]string{"a", "b", "c", "d"}

	fmt.Println(a1)

	fmt.Println(a2)

	fmt.Println(a3)

	fmt.Println(a4)

}
