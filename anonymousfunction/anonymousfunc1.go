package main

import "fmt"

func main() {

	func(num1, num2, num3 int) int {

		total := num3 * (num1 + num2)

		fmt.Println("Total:", total)

		return total

	}(5, 3, 2) // this part is func call

}
