package main

import (
	"fmt"
)

func main() {
	var x, y, z int = 10, 20, 30 //declqaring multiple values at a time

	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	//variables inside a block
	var (
		id      int
		name    string = "johnny"
		country string = "uk"
	)
	fmt.Println(id)
	fmt.Println(name)
	fmt.Println(country)

}
