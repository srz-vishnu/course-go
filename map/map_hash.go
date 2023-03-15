package main

import (
	"fmt"
)

func main() {

	z := make(map[string]string)

	z["First"] = "Dan"

	z["Second"] = "Isaac"

	z["Third"] = "Jane"

	x := z

	y := z

	fmt.Println("Before: ")

	fmt.Println(x)

	fmt.Println(y)

	fmt.Println(z)

	x["First"] = "YYYYYYYYYYYYYYYYYYYYYYY" //this affect in y and z map also, bcoz map references hash tables

	fmt.Println("After: ")

	fmt.Println(x)

	fmt.Println(y)

	fmt.Println(z)

}
