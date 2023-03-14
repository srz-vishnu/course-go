package main

import (
	"fmt"
)

func main() {

	var num = [4]int{7, 8, 9, 10}

	num2 := [6]int{11, 12, 13, 14, 15, 16}

	num3 := [...]int{1, 2, 3, 4, 5, 6} // go decides the length of array thats what 3dots stands for

	name := [3]string{"john", "vicky", "sam"}

	fmt.Println(num)

	fmt.Println(num2)

	fmt.Println(num3)

	fmt.Println(name)

}
