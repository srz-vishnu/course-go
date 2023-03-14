package main

import (
	"fmt"
)

// ques to print its value n datatype using one output method

func main() {

	var a = 30

	var b = "Hello"

	var c = true

	fmt.Printf("%v'is' %T\n", a, a) //%v is 30 n %T is datatype

	fmt.Printf("%v'is'%T\n", b, b)

	fmt.Printf("%v'is'%T\n", c, c)

}
