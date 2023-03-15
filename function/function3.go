package main

import "fmt"

func num(x float64) (a, b, c float64) {
	a = x - 2.25
	b = x - 0.25
	c = x - 1.1
	return a, b, c
}

func main() {

	// m, n, p := num(10)  // or we can use that code also
	// println(m, n, p)

	var i float64
	fmt.Println("Enter a number: ")
	fmt.Scan(&i)

	m, n, p := num(i)
	fmt.Println(m, n, p)

}
