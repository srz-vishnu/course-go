package main

import (
	"fmt"
)

type User struct {
	firstName string
	lastName  string
	age       int
	fee       int
}

func main() {

	var user1 User //zero valued struct

	fmt.Println("First Name:", user1.firstName)

	fmt.Println("Last Name:", user1.lastName)

	fmt.Println("Age:", user1.age)

	fmt.Println("Fee:", user1.fee)

}
