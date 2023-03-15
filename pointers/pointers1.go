package main

import (
	"fmt"
)

func main() {

	var i int

	var j *int

	i = 800

	j = &i

	fmt.Println("Before: ", i)

	*j = 1400

	fmt.Println("After: ", i)

}
