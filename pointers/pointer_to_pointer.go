package main

import (
	"fmt"
)

func main() {

	var m = 58.3

	var p = &m

	var pp = &p

	fmt.Println("m: ", m)

	fmt.Println("p: ", p)

	fmt.Println("&m: ", &m)

	fmt.Println("pp: ", pp)

	fmt.Println("*p: ", *p)

	fmt.Println("**pp: ", **pp)

}
