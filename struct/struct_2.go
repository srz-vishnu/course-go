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

	user1 := User{

		age:       55,
		firstName: "Dan",
		fee:       8000,
		lastName:  "Andrew",
	}

	user2 := User{"Kate", "Vincent", 41, 5500}

	fmt.Println("User 1:", user1)
	fmt.Println("User 2:", user2)

}
