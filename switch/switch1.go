package main

import "fmt"

func main() {

	name := "tim"

	switch name {

	case "Peter":

		fmt.Println("Hi Peter")

	case "Tim":

		fmt.Println("Hi Tim")

	case "Vic":

		fmt.Println("Hi Vic")

	case "James":

		fmt.Println("Hi James")

	case "Anne":

		fmt.Println("Hi Anne")

	default:
		fmt.Println("unknown")

	}

}
