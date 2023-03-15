package main

import (
	"fmt"
)

func main() {

	t := func() func(num1, num2, num3 int) {

		w := func(num1, num2, num3 int) {

			total := num3 * (num1 + num2)

			fmt.Println("Total:", total)

		}

		return w

	}

	z := t()

	z(2, 3, 4)

}
