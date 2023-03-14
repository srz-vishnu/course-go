package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 28

	n := 12.5
	m := 4.2

	result := n / m

	fmt.Printf("%g / %g = %g\n", n, m, result) //%g when its float
	//fmt.Printf("m / n = %g\n", m/n)

	fmt.Printf("a + b = %d\n", a+b)

	fmt.Printf("a - b = %d\n", a-b)

	fmt.Printf("a * b = %d\n", a*b)

	fmt.Printf("a / b = %d\n", a/b)

	fmt.Printf("a mod b = %d\n", a%b)

	fmt.Printf("a = %d\n", a)

	a++

	fmt.Printf("a++ = %d\n", a)

	fmt.Printf("b = %d\n", b)

	b--

	fmt.Printf("b-- = %d\n", b)
}
