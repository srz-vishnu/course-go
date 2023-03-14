package main

import "fmt"

func main() {

	x := 5

	y := 10

	z := 15

	fmt.Println(x < y && x > z)

	fmt.Println(x < y || x > z)

	fmt.Println(!(x < y && x > z))

}
