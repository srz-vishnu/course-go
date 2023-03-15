package main

import "fmt"

func main() {

	age := 17

	if age < 18 {

		fmt.Println("Your age is below the required age")

	} else if age >= 18 && age <= 65 {

		fmt.Println("Your are eligible")

	} else if age > 65 && age < 70 {

		fmt.Println("Under consideration")

	} else {

		fmt.Println("You are too old to apply")

	}

}
