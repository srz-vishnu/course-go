package main

import (
	"fmt"
)

func main() {

	var x int = 6125

	var y *int = &x

	fmt.Println(x)

	fmt.Println(y)

	fmt.Println(*y)

}
