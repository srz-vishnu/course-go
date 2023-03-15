package main

import "fmt"

func main() {

	age := 10

	if age > 15 {

		fmt.Println("You are more than 15")

		if age >= 18 {

			fmt.Println("You are eligible")

		}

	} else {

		fmt.Println("You are below the age")

	}

}
