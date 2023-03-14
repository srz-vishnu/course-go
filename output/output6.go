package main

import (
	"fmt"
)

//ques is to print values of m and n in hexadecimal and binary

func main() {

	n := 12

	m := 25

	fmt.Printf("%v is : %b in binary\n", n, n) //%b to denote binary and %X to hexadecimal

	fmt.Printf("%v is : %X in hexadecimal\n", m, m)

}
