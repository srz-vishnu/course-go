package main

import "fmt"

func main() {

	month := "May"

	switch month {

	case "January", "March", "May", "July", "August", "October", "December":

		fmt.Println("There are 31 days in", month)

	case "April", "June", "September", "November":

		fmt.Println("There are 30 days in", month)

	case "February":

		fmt.Println("There are 28 days in %s\n", month)

	default:

		fmt.Println("Invalid name of month", month)

	}

}
