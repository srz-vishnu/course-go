package main

import (
	"fmt"
)

func main() {
	const X, Y, Z int = 10, 20, 30 //declqaring multiple values at a time

	fmt.Println("x:", X)
	fmt.Println("y:", Y)
	fmt.Println("z:", Z)

	// const declartion using a block
	const (
		ID      int    = 65
		NAME    string = "johnny"
		COUNTRY string = "uk"
	)
	fmt.Println(ID)
	fmt.Println(NAME)
	fmt.Println(COUNTRY)

}
