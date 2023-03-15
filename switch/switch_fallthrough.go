package main

import "fmt"

func main() {

	name := "Peter"

	switch name {

	case "Peter":

		fmt.Println("Hi Peter") //peter had match so there we can use fallthrough..fallthrough helps to print the next statement..so it prints peter n tim
		fallthrough

	case "Tim":

		fmt.Println("Hi Tim") ///we can use fallthrough here just bcoz of fallthrough used on the peter statement, bcoz of fallthrough here vic will also print
		fallthrough

	case "Vic":

		fmt.Println("Hi Vic")

	case "James":

		fmt.Println("Hi James") //fallthrough wont wrk here bcoz we dont have a match and we dont have support of matched top statement
		fallthrough

	case "Anne":

		fmt.Println("Hi Anne")

	default:
		fmt.Println("unknown")

	}

}
