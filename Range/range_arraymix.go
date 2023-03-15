package main

import "fmt"

func main() {

	names := []string{"David", "Jane"}

	items := []string{"apple", "orange", "banana"}

	for _, name := range names {

		fmt.Printf("%s:\n", name)

		for _, item := range items { // 1st david edkum then inner loop il kerum, all values excute cheyum, then exit avum so outer loop il povum name marum again inner loop

			fmt.Println(item)

		}

		fmt.Println()

	}

}
