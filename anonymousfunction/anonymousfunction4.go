package main

import (
	"fmt"

	"math"
)

var (
	c = func(r float64) float64 {

		return 2 * math.Pi * r

	}
)

func main() {

	fmt.Println(c(10))

}
